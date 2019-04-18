#! /bin/bash

set -euo pipefail

function log() { printf "%4d: %s\n" $SECONDS "$*"        >&2 ; }
function LOG_D() { printf "%4d [DEBUG] %s\n" $SECONDS "$*"        >&2 ; }
function die() { printf "%4d: ERROR: %s\n" $SECONDS "$*" ; exit 1 ; }

mkdir -p benchmark

TESTDIR=$PWD
WIPE_DATADIR=false
OBJECTBOX_DB_DIR=benchmark-test


export GIT_TERMINAL_PROMPT=1
TMPFS_MOUNTPOINT=/ramtmp
TMPDIR=$(realpath benchmark)
N=1
export GOPATH=$HOME/go

function showHelp() {
    echo "Objectbox's $(basename $0) help:"
    echo "  -h|--help             - This help"
    echo "  -r|--repetitions      - Set number of repetitions (default: $N)"
    echo "  -c|--item-count       - Set the default item count (default: ${EDGEX_TEST_COUNT:-})"
    echo "  -k|--ignore-errors    - Do not stop on errors"
    echo "  --ram=yes|no|auto     - Use a RAM filesystem for data"
    echo "                          - no (default): Use \$TMPDIR: $TMPDIR"
    echo "                          - yes: Mount it if necessary (needs sudo or an fstab entry)"
    echo "                          - auto: Use a tmpfs at ${TMPFS_MOUNTPOINT} if available"
    echo
    echo "  --only-env            - Only prepare environment and exit without running tests"
    echo "  --no-build            - Don't rebuild tests"
    echo "  -g                    - Set CPU performance to maximum" 
    echo
    echo "  -a|--all-backends     - Enable all backends"
    echo "  -O|--objectbox              - Enable Objectbox test"
    echo "  -R|--redis            - Enable Redis test"
    echo "  -M|--mongo            - Enable Mongo test"
    
    echo "  --remove-gopath       - Wipe the GOPATH"
    echo "  --remove-gopath=even-if-std - Wipe the GOPATH even if its the standard one (~/go)"
    echo
    cat <<END
 * The software source must have been deployed to the current GOPATH:  ${GOPATH}

 * If software is to be downloaded from git, you may need to add "-A" (ForwardAgent=yes) to your ssh line.
 
 * It is possible -and even healthy- to use an alternate GOPATH. 
END
}

while [ $# -gt 0 ]
do
    case $1 in
    -h|--help)
        showHelp
        exit 0
        ;;
    -r|--repetitions)
        N=$2
        log "Setting N==$N"
        shift
        ;;
    -c|--item-count)
        export EDGEX_TEST_COUNT=$2
        log "EDGEX_TEST_COUNT=$EDGEX_TEST_COUNT"
        shift
        ;;
    -k|--ignore-errors)
        IGNORE_ERRORS=true
        log "IGNORE_ERRORS=true"
        ;;
    -O|--objectbox)
        ENABLE_OBJECTBOX=true
        SOME_TEST_ENABLED=true
        ;;
    -R|--redis)
        ENABLE_REDIS=true
        SOME_TEST_ENABLED=true
        ;;
    -M|--mongo)
        ENABLE_MONGO=true
        SOME_TEST_ENABLED=true
        ;;
    --ram=*)
        TMPMODE=${1#*=}
        echo TMPMODE==$TMPMODE
        ;;
    --only-env)
        ONLY_PREPARE_ENVIRONMENT=true
        ;;
    --gopath)
        export GOPATH=$2
        shift
        ;;
    --remove-gopath)
        if [ "$GOPATH" == $HOME/go ]
        then
            die "I don't dare to remove your standard GOPATH $GOPATH. Use --remove-gopath=even-if-std"
        fi
        REMOVE_GOPATH=true
        ;;
    --remove-gopath=even-if-std)
        REMOVE_GOPATH=true
        ;;
    --no-build)
        NO_BUILD=true
        ;;
    -g)
        CPU_GOVERNOR=$2
        shift
        ;;
    *)
        echo "Bad option: $1" >&2
        echo
        showHelp
        exit 1
    esac
    shift
done

if [[ ${CPU_GOVERNOR:-} ]]
then
    echo "Setting CPU clock governor to ${CPU_GOVERNOR}"
    SAVED_GOVERNOR="$(cpufreq-info -p |awk '{print $NF}')"
    sudo cpufreq-set -g ${CPU_GOVERNOR}    
    trap "echo Resetting CPU governor $SAVED_GOVERNOR... ; set -x ; sudo cpufreq-set -g $SAVED_GOVERNOR" EXIT
fi

if ${REMOVE_GOPATH:-false}
then
    echo  -n "Removing GOPATH $GOPATH..."
    rm -rf $GOPATH
    echo " done."
fi

## class Sectors {

function initializeDiskStats() {
	PART=$(findmnt -n -o SOURCE --target "$1" | cut -f3 -d/ |sed 's/p[0-9]$//')
	DEV=$(for f in /sys/block/* 
			do 
				if echo "$PART" | grep -q ${f##*/} ; then 
					echo $f ; break ;
				fi			
		 done
		 )
}

function sectorsWrittenInDisk() { 
	if [[ "$DEV" ]]
	then
		awk '{print $(7+1)}'  $DEV/stat
	else
		return 1
	fi
}

getResetWrittenSectors() {
	if sync && WRITTEN_SECTORS2=$(sectorsWrittenInDisk) ; then
		
		log "$1: Written sectors: " $(( $WRITTEN_SECTORS2 - $WRITTEN_SECTORS ))
		echo "$ENGINE $(( $WRITTEN_SECTORS2 - $WRITTEN_SECTORS ))" |tee -a $DATADIR/$ENGINE.iostats
		WRITTEN_SECTORS=$WRITTEN_SECTORS2
	fi 
}

## } // sectors

OBX_PKG=github.com/objectbox/objectbox-go

case ${TMPMODE:-no} in
        on|yes)
            if ! [ -d ${TMPFS_MOUNTPOINT} ] ; then
                log "Attempt to create ${TMPFS_MOUNTPOINT}: sudo mkdir -v ${TMPFS_MOUNTPOINT}"
                sudo mkdir -v ${TMPFS_MOUNTPOINT}
            fi
            if ! df ${TMPFS_MOUNTPOINT} | grep -q tmpfs ; then
                log "Attempting to mount a tmpfs on ${TMPFS_MOUNTPOINT}: sudo mount -t tmpfs tmpfs ${TMPFS_MOUNTPOINT}"
                sudo mount -t tmpfs tmpfs ${TMPFS_MOUNTPOINT}
            fi
            log "Using RAM disk at ${TMPFS_MOUNTPOINT}"
            TMPDIR=${TMPFS_MOUNTPOINT}
            ;;
        off|no)

            ;;
        auto)
            if df ${TMPFS_MOUNTPOINT} | grep -q tmpfs ; then
                log "Detected RAM disk at ${TMPFS_MOUNTPOINT}"
                TMPDIR=${TMPFS_MOUNTPOINT}
            fi
            ;;
        *)
        die "Bad --ram option. Options: yes|no|auto"
esac

initializeDiskStats $TMPDIR

if ! ${SOME_TEST_ENABLED:-false}
then
    log "Enable all tests"
    ENABLE_OBJECTBOX=true
    ENABLE_MONGO=true
    ENABLE_REDIS=true
fi
: ${ENABLE_OBJECTBOX:=false}
: ${ENABLE_REDIS:=false}
: ${ENABLE_MONGO:=false}

TTY_OR_EMPTY="$(tty >& /dev/null && tty)" || LOG_D "No connected tty."


TMPF=${TMPDIR:=/tmp}/edgex-benchmark.tmp
TIMESTMPFILE=${TMPDIR:-/tmp}/edgex-benchmark-times.tmp

export PATH=$PATH:/sbin:/usr/sbin ## ldconfig
##########

function requireBinaries() {
    for bin in "$@"; do
        if ! which ${bin} >& /dev/null; then
            die "'${bin}' not found in the PATH"
        fi
    done
}

requireBinaries go

###########

log "Building binaries..."

if ${ONLY_PREPARE_ENVIRONMENT:-false}; then
    exit 0
fi

OBJECTBOX_TESTBIN=$TMPDIR/github_com_edgexfoundry_edgex_go_internal_pkg_db_objectbox
MONGO_TESTBIN=$TMPDIR/github_com_edgexfoundry_edgex_go_internal_pkg_db_mongo
REDIS_TESTBIN=$TMPDIR/github_com_edgexfoundry_edgex_go_internal_pkg_db_redis
REDIS_CONF=$TESTDIR/redis.conf

function addDependencyFile() {
    DEPENDENCIES="${DEPENDENCIES:-} $*"
}

if ${ENABLE_OBJECTBOX}; then 
    addDependencyFile $OBJECTBOX_TESTBIN
    addDependencyFile $(ldconfig -p | awk '/libobjectbox.so/ {print $NF}')
fi

if ${ENABLE_MONGO}; then 
    addDependencyFile $MONGO_TESTBIN
fi

if ${ENABLE_REDIS}; then
    addDependencyFile $REDIS_TESTBIN
    if [[ -f $REDIS_CONF ]]; then
        addDependencyFile $REDIS_CONF
    else
        REDIS_CONF=
    fi
fi

if ! ${NO_BUILD:-false} ; then
    ${ENABLE_OBJECTBOX} && {
        log "Building ObjectBox benchmark, binary = $OBJECTBOX_TESTBIN..."
        go build -o $OBJECTBOX_TESTBIN ./internal/pkg/db/objectbox/benchmark
    }
    ${ENABLE_MONGO} && {
        log "Building Mongo benchmark, binary = $MONGO_TESTBIN..."
        requireBinaries mongod
        go build -o $MONGO_TESTBIN ./internal/pkg/db/mongo/benchmark
        
    }
    ${ENABLE_REDIS} && {
        log "Building Redis benchmark, binary = $REDIS_TESTBIN..."
        requireBinaries redis-server
        go build -o $REDIS_TESTBIN ./internal/pkg/db/redis/benchmark

    }
fi

log "Binaries built."

## *********************************
# Calculate hash of used binaries (tests + libobjectbox.so)

NL=$'\n'
TAB=$'\t'

function addConfigCmd() {
    local key="$1"
    shift
    EXTRA_STATE="${EXTRA_STATE:-}${key}${TAB}$("$@" | tr '\n' ' ')"$NL
}

function addVar() {
    local varname=$1
    EXTRA_STATE="${EXTRA_STATE:-}${varname}${TAB}${!varname:-<unassigned>}"${NL}
}

addConfigCmd  dataPartition bash -c "df $TMPDIR | awk 'NR==2 {print \$1}' "
addVar EDGEX_TEST_COUNT 
addConfigCmd "go version" go version
#addConfigCmd "cpufreq-info -p"  cpufreq-info  -p
addConfigCmd prio bash -c 'chrt -v -p $$ | sed "s/$$/PID/g"'
if ${ENABLE_MONGO}; then
    addConfigCmd mongod-version mongod --version
fi
if ${ENABLE_REDIS}; then
    addConfigCmd redis-server-version redis-server --version
fi
#addConfigCmd hdType cat /sys/class/block/${PART}/device/model || ## If it doesn't work for the tmp dir, we register it for the rootfs, just in case it's the same
#   addConfigCmd rootHdType bash -c 'cat /sys/class/block/$(findmnt -n -o SOURCE --target / | cut -f3 -d/ |sed "s/p[0-9]$//")/device/model'
   
addConfigCmd uname-a uname -a

cc -o $TMPDIR/objectbox_version_core_string -lobjectbox -xc - <<EOP
#include <stdio.h>
extern const char* obx_version_core_string(void);  //header file is probably not available here
int main() { puts(obx_version_core_string()); }
EOP
addConfigCmd objectbox-version $TMPDIR/objectbox_version_core_string
rm -f $TMPDIR/objectbox_version_core_string

RUNHASH=$(
        (
        xargs -r md5sum <<< "${DEPENDENCIES:-}" |  
            awk '{print gensub(/.*\//, "", "g", $2) " " $1 ; }'    ## basename file and swap with md5sum
        echo $EXTRA_STATE
        ) | md5sum - | cut -f1 -d ' ')

log "RUNHASH = $RUNHASH"


DATADIR=$TMPDIR/$RUNHASH
if mkdir -vp $DATADIR
then
    (md5sum ${DEPENDENCIES:-}
        echo "$EXTRA_STATE"
        )  | tee  $DATADIR/config.txt
fi

lscpu > $DATADIR/cpu.txt
top -bn1 > $DATADIR/top.txt
echo "Load averages:"
cat /proc/loadavg | tee $DATADIR/loadavg.txt

function nextid() {
    LAST_ID_FILE=$DATADIR/lastid
    LOCKFILE=$DATADIR/lockfile
     (
         flock -n 9 || exit 1
         LAST_ID=$(cat ${LAST_ID_FILE}  || echo 0)
         NEXT_ID=$(( $LAST_ID + 1 )) 
         echo $NEXT_ID > ${LAST_ID_FILE}
         echo $NEXT_ID
       ) 9>${LOCKFILE}
}

function execute() {
    ENGINE=$1
    EXEC_LOG_FILE=$DATADIR/$ENGINE-out.${EXECUTION_SEQ_ID}.txt
    log "Executing $1, repetition $i... writing to ${EXEC_LOG_FILE}"
    shift
    STATUS=0
	WRITTEN_SECTORS=$(sync; sectorsWrittenInDisk) || log "Disk stats are not available for partition $PART"
    \time -v -o $TIMESTMPFILE -- "$@"  |& tee ${EXEC_LOG_FILE} $TTY_OR_EMPTY > $TMPF || STATUS=${PIPESTATUS[0]}
    getResetWrittenSectors $ENGINE
    case $STATUS in
    130)
        die SIGINT
        ;;
    143)
        die SIGTERM
        ;;
    0)
      cat $TIMESTMPFILE >> $DATADIR/$ENGINE.times
        ( echo -ne "$ENGINE\t$(date '+%Y-%m-%d %R')\t$(< $TMPF sed -n  's/.*(\([0-9.]\+\) iterations per second.*/\1/p' | tr '.\n' ',\t' )"
            < $TMPF sed -n 's/^Making changes durable: \(.*\)ms$/\t\1/p' | tr '.\n' ',\t'
            echo
        ) | tee  $TTY_OR_EMPTY >> ${DATADIR}/${ENGINE}.csv
            ;;
    *)
        log "$ENGINE failed $STATUS: "
        cat $TMPF

	    if ${ONCE_OK:-false}
	    then
		      log "But this was successful once"
	    fi

	    ERRORS=true

        if ! ${IGNORE_ERRORS:-false}
        then
            return 1
        fi
    esac
}



function run_objectbox() {
    set -e
    cd $TMPDIR
    if ${WIPE_DATADIR}; then
        log "Remove objectbox/..."
        rm -rf ${OBJECTBOX_DB_DIR}
    fi
    log "Starting objectbox test... "
    execute objectbox $OBJECTBOX_TESTBIN
    du -sm ${OBJECTBOX_DB_DIR} >> $DATADIR/$ENGINE.datasize || die "No objectbox directory ${OBJECTBOX_DB_DIR} found"
    cd -
}

function run_redis() {
    DBDIR=$TMPDIR/redisdb
    
    du -h $DBDIR || true
    if ${WIPE_DATADIR}
    then
        rm -fvr $DBDIR/*
    fi
    mkdir $DBDIR || true
    cd $DBDIR
    \time -vao "$DATADIR/redis-server.times" -- redis-server $REDIS_CONF >& $DATADIR/redis-server.${EXECUTION_SEQ_ID}.log & REDIS_PID=$!
    cd -
    trap "pgrep -P $REDIS_PID |xargs -r kill " EXIT
    waitForTcpOpen 6379 60
    log "Now starting redis test..."
    WRITTEN_SECTORS=$(sync ; sectorsWrittenInDisk) || log "Disk stats are not available for partition $PART"
    execute redis $REDIS_TESTBIN
    pgrep -P $REDIS_PID |xargs -r kill || log "Redis not there"
    trap EXIT
    wait $REDIS_PID   # We kill the actual process but wait for the "time" process, which is our shell's child
    
    du -sm $DBDIR | tee $DATADIR/$ENGINE.datasize
}

function waitForTcpOpen() {
    log "Waiting for port $1..."
    n=0
    while ! nc -z localhost $1
    do
        n=$(( $n + 1 ))
        if [ $n -eq ${2:-10} ]
        then
            die "Port $1 didn't get open"
        fi

        log "Waiting for open $1 to be open ...  $n" >&2
        sleep 1
    done
}

function run_mongo() {
    #DBD="$(sed -n 's/dbpath=//p' mongodb.conf)"
    DBPATH=$TMPDIR/mongodb-$USER
    DBPATH_TARBALL=$TMPDIR/mongodb-$USER.tar

    rm -rf $DBPATH
    mkdir -vp $DBPATH
    if ${WIPE_DATADIR} && { [ ! -f $DBPATH_TARBALL ] || ! tar xavf $DBPATH_TARBALL -C $DBPATH ; }
    then
        log "Initializing Mongo, only FIRST TIME ... ===================="
        rm -f $DBPATH_TARBALL
        \time -vao mongo-setup.times -- mongod   --smallfiles --unixSocketPrefix=/tmp  --dbpath=$DBPATH & MONGO_PID=$!
        trap "kill $MONGO_PID" EXIT
        waitForTcpOpen 27017 60
        pgrep -P $MONGO_PID |xargs -r kill || log "Mongo server not there"
        cd $DBPATH
        log "Waiting for Mongo to finish..."
        wait $MONGO_PID
        log "Done."

        tar cavf $DBPATH_TARBALL.tmp  .
        mv  $DBPATH_TARBALL.tmp $DBPATH_TARBALL -v

        cd -
    fi

    \time -vao $DATADIR/mongo-server.times -- mongod   --smallfiles --unixSocketPrefix=/tmp  --dbpath=$DBPATH >& $DATADIR/mongo-server.${EXECUTION_SEQ_ID}.log & MONGO_PID=$!
    trap "kill $MONGO_PID" EXIT
    waitForTcpOpen 27017 60
    if ! kill -0 $MONGO_PID ; then
        die "Mongo died."
    fi
    while ! mongo <<< "show dbs"
    do
        echo Mongo is a bit overwhelmed, waiting...
        sleep 1
        if ! kill -0 $MONGO_PID ; then
            die "Mongo died."
        fi
    done
    sleep 2
    echo Now starting mongo test... >&2
    execute mongo $MONGO_TESTBIN
    log "Test finished, killing Mongo..."
    pgrep -P $MONGO_PID | xargs -r kill || log "Mongo server not there"

    wait $MONGO_PID

    du -sm $DBPATH > $DATADIR/$ENGINE.datasize
    trap EXIT
}



for i in $(seq 1 $N)
 do
    EXECUTION_SEQ_ID=$(nextid)
    log "Iteration #$i. Seq id: ${EXECUTION_SEQ_ID}"

    ${ENABLE_OBJECTBOX}   && run_objectbox
    ${ENABLE_REDIS} && run_redis
    ${ENABLE_MONGO} && run_mongo

    if ! ${ERRORS:-false}
    then
        ONCE_OK=true
    fi
done

echo "== Success in $SECONDS seconds. Hash == $RUNHASH :)" >&2


## Evaluation

cd $DATADIR
if [ -f objectbox.csv ]; then 
    logs="$(ls -f {redis,mongo}.csv 2> /dev/null | xargs -r echo)" || log  "not all engines available"
    if [[ $logs ]]; then
        for col in {2..9}; do
            echo "Factors $logs, col $col:  $(tail -qn 1 objectbox.csv $logs |awk -v col=$(( $col  + 2 )) 'NR==1 { ref=$col; next; } $col && NR >1 { printf ref / $col " "; } ' )"|| true
        done
    else
        log "Only objectbox available"
    fi
    cat *.csv
fi

