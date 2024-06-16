if [ "$#" -eq 0 ]; then
    echo "Illegal number of parameters"
    exit 1
fi

COMMAND_GROUPS=""

if [ -n "$2" ]; then
    COMMAND_GROUPS="--groups $2"
fi

firebase appdistribution:distribute --app $1 $COMMAND_GROUPS