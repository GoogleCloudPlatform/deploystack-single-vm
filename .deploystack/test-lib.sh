CYAN='\033[0;36m'
BCYAN='\033[1;36m'
NC='\033[0m' # No Color
DIVIDER=$(printf %"$(tput cols)"s | tr " " "*")
DIVIDER+="\n"

function get_project_id() {
    local __resultvar=$1
    VALUE=$(gcloud config get-value project | xargs)
    eval $__resultvar="'$VALUE'"
}

function get_project_number() {
    local __resultvar=$1
    local PRO=$2
    VALUE=$(gcloud projects list --filter="project_id=$PRO" --format="value(PROJECT_NUMBER)" | xargs)
    eval $__resultvar="'$VALUE'"
}

# DISPLAY HELPERS
function section_open() {
    section_description=$1
    printf "$DIVIDER"
    printf "${CYAN}$section_description${NC} \n"
    printf "$DIVIDER"
}

function section_close() {
    printf "$DIVIDER"
    printf "${CYAN}$section_description ${BCYAN}- done${NC}\n"
    printf "\n\n"
}

function evalTest() {
    local command=$1
    local expected=$2

    local ERR=""
    got=$(eval $command 2>errFile)
    ERR=$(<errFile)

    if [ ${#ERR} -gt 0 ]; then
        if [ "$expected" = "EXPECTERROR" ]; then
            printf "Expected Error thrown \n"
            return
        fi

        printf "Halting - error: '$ERR'  \n"
        exit 1
    fi

    if [ "$got" != "$expected" ]; then
        printf "Halting: '$got'  \n"
        exit 1
    fi

    printf "$expected is ok\n"
}


function generateProject(){
    local __resultvar=$1
    local __STACKSUFFIX=$2
    local __RANDOMSUFFIX=$(
        LC_ALL=C tr -dc 'a-z0-9' </dev/urandom | head -c 8
        echo
    )
    local __DATELABEL=$(date +%F)
    local VALUE=ds-test-$__STACKSUFFIX-$__RANDOMSUFFIX
    local __BA=$(gcloud beta billing accounts list --format="value(ACCOUNT_ID)" --limit=1 | xargs)
   
    gcloud projects create $VALUE --labels="deploystack-disposable-test-project=$__DATELABEL" --folder="155265971980"
    gcloud beta billing projects link $VALUE --billing-account=$__BA
    eval $__resultvar="'$VALUE'"
}
