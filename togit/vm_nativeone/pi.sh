# Check and Usage
function Usage {
    echo "Usage:    ./reg-enr-peer.sh   'CA_ADDRESS'   'ORGANIZATION_NAME(LOWERCASE)'   'PEER_NAME(LOWERCASE)'"
    echo "Usage:    ./reg-enr-peer.sh   192.168.50.50   one   one-peer1"
}

CA_ADDRESS=$1
if [ -z $CA_ADDRESS ];
then
    Use
    echo "Provide the CA Address"
    exit 0
fi

ORGANIZATION_NAME=$2
if [ -z $ORGANIZATION_NAME ];
then
    Use
    echo "Provide the Organization name"
    exit 0 
fi

PEER_NAME=$3
if [ -z $PEER_NAME ];
then
    Use
    echo "Provide the Peer name"
    exit 0 
fi



Port=7054

# Set Path for Organization Admin MSP
export FABRIC_CA_CLIENT_HOME=~/fabric/CA/clients/$ORGANIZATION_NAME/admin

# Register Peer
fabric-ca-client register --id.type peer --id.name $PEER_NAME --id.secret pw --id.affiliation $ORGANIZATION_NAME

# Set Path for Peer MSP
export FABRIC_CA_CLIENT_HOME=~/fabric/CA/clients/$ORGANIZATION_NAME/$PEER_NAME

# Enroll Peer
fabric-ca-client enroll -u http://$PEER_NAME:pw@$CA_ADDRESS:$Port

# Copy the requied Admincerts for MSP
mkdir -p $FABRIC_CA_CLIENT_HOME/msp/admincerts
cp ~/fabric/CA/clients/$ORGANIZATION_NAME/admin/msp/signcerts/* $FABRIC_CA_CLIENT_HOME/msp/admincerts 


