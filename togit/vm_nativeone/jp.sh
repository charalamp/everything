
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
if [ -z $ORDERER_ADDRESS ];
then
    Use
    echo "Provide the Orderer Address"
exit 0
fi



# Set the variables(core.yaml)
export CORE_PEER_MSPCONFIGPATH=~/fabric/CA/clients/$ORGANIZATION_NAME/admin/msp

export CORE_PEER_LOCALMSPID=$MSP_ID
export CORE_PEER_ID=$PEER_ID

export CORE_PEER_LISTENADDRESS=$PEER_ADDRESS:7051
export CORE_PEER_ADDRESS=$PEER_ADDRESS:7051
export CORE_PEER_CHAINCODELISTENADDRESS=$PEER_ADDRESS:7052
export CORE_PEER_EVENTS_ADDRESS=$PEER_ADDRESS:7053

export CORE_PEER_GOSSIP_BOOTSTRAP=$PEER_ADDRESS:7051
export CORE_PEER_GOSSIP_ENDPOINT=$PEER_ADDRESS:7051
export CORE_PEER_GOSSIP_EXTERNALENDPOINT=$PEER_ADDRESS:7051

export CORE_PEER_FILESYSTEMPATH=~/fabric/artifacts/ledger

export FABRIC_CFG_PATH=~/fabric/config



# Fetch the Genesis block(Application channel) from the Orderer
peer channel fetch config onetwochannel-genesis.block -o $ORDERER_ADDRESS:7050 -c onetwochannel

# Join Application Channel
peer channel join -o $ORDERER_ADDRESS:7050 -b  onetwochannel-genesis.block


