function Usage {
    echo "Usage:    ./instantiatecc.sh    [PEER_ADDRESS]   [ORGANIZATION_NAME(LOWERCASE)]   [PEER_ID]   [MSP_ID(CONFIGTX.YAML)]   [ORDERER_ADDRESS]"
    echo "Usage:    ./instantiatecc.sh   192.168.50.50   one   one-peer1   OneMSP   192.168.50.51"   
}

PEER_ADDRESS=$1
if [ -z $PEER_ADDRESS ];
then
    Use
    echo "Provide the Peer Address"
    exit 0
fi

ORGANIZATION_NAME=$2
if [ -z $ORGANIZATION_NAME ];
then
    Use 
    echo "Provide the Organization name"
    exit 0 
fi

PEER_ID=$3
if [ -z $PEER_ID ];
then
    Use
    echo "Provide the Peer name"
    exit 0 
fi

MSP_ID=$4
if [ -z $MSP_ID ];
then
    Use
    echo "Provide the MspID name"
    exit 0 
fi

ORDERER_ADDRESS=$5
if [ -z $MSP_ID ];
then
    Use
    echo "Provide the Orderer Address"
exit 0
fi



# Set the variables(core.yaml)
export CORE_PEER_MSPCONFIGPATH=~/fabric/CA/clients/one/admin/msp

export CORE_PEER_LOCALMSPID=OneMSP
export CORE_PEER_ID=one-peer1

export CORE_PEER_LISTENADDRESS=$PEER_ADDRESS:7051
export CORE_PEER_ADDRESS=$PEER_ADDRESS:7051
export CORE_PEER_CHAINCODELISTENADDRESS=$PEER_ADDRESS:7052
export CORE_PEER_EVENTS_ADDRESS=$PEER_ADDRESS:7053

export CORE_PEER_GOSSIP_BOOTSTRAP=$PEER_ADDRESS:7051
export CORE_PEER_GOSSIP_ENDPOINT=$PEER_ADDRESS:7051
export CORE_PEER_GOSSIP_EXTERNALENDPOINT=$PEER_ADDRESS:7051

export CORE_PEER_FILESYSTEMPATH=~/fabric/artifacts/ledger

export FABRIC_CFG_PATH=~/fabric/config

#CC_CONSTRUCTOR='{"Args":["init","a", "100", "b","200"]}'

# Instantiate Chaincode
#peer chaincode instantiate -C onetwochannel -n marbles -v 1.0 -c '{"Args":["init","a", "100", "b","200"]}'  -o  $ORDERER_ADDRESS:7050 

#peer chaincode invoke -C onetwochannel -n marbles -c '{"Args":["invoke","a","b","10"]}'  -o  $ORDERER_ADDRESS:7050


peer chaincode query -C onetwochannel -n marbles -c '{"Args":["query","a"]}'   -o  $ORDERER_ADDRESS:7050
