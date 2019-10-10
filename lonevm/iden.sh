########### Check

function Use {
    echo "Use:    ./identities.sh   [CA_ADDRESS]   [ORGANIZATION_NAME(LOWERCASE)]"
    echo "Use:    ./identities.sh   192.168.50.50   one"
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


Port=7054

########### Start CA

# Set Fabric CA Server Home Folder
export FABRIC_CA_SERVER_HOME=/home/vagrant/fabric/CA/server

# Init CA
fabric-ca-server init

# Start CA
fabric-ca-server start 2> /home/vagrant/fabric/ca.log &



sleep 3s




echo '########### Enroll CA Admin'



# Set Path for CA Registar/Admin MSP
export FABRIC_CA_CLIENT_HOME=~/fabric/CA/clients/admin 

# Enroll CA Registar/Admin
fabric-ca-client enroll -u http://admin:pw@$CA_ADDRESS:$Port

sleep 3s



echo '########### Register Organization Admin'


# Set Path for CA Registrar/Admin MSP 
export FABRIC_CA_CLIENT_HOME=~/fabric/CA/clients/admin

# Register Organization Admin
fabric-ca-client register --id.type admin --id.name $ORGANIZATION_NAME-admin --id.secret pw --id.affiliation $ORGANIZATION_NAME --id.attrs '"hf.Registrar.Roles=peer,orderer,user,client","hf.AffiliationMgr=true","hf.Revoker=true"'  --id.attrs 'one=true:ecert' 


sleep 3s






echo '########### Enroll Organization Admin'


# Set Path for Organization Admin MSP
export FABRIC_CA_CLIENT_HOME=~/fabric/CA/clients/$ORGANIZATION_NAME/admin

# Enroll Organization Admin
fabric-ca-client enroll -u http://$ORGANIZATION_NAME-admin:pw@$CA_ADDRESS:$Port

# Copy the required admincerts for MSP
mkdir -p $FABRIC_CA_CLIENT_HOME/msp/admincerts
cp ~/fabric/CA/clients/admin/msp/signcerts/*  $FABRIC_CA_CLIENT_HOME/msp/admincerts

sleep 3s







echo '########### Set Organization MSP'


# Set Path for Organization MSP
export FABRIC_CA_CLIENT_HOME=~/fabric/CA/clients/$ORGANIZATION_NAME

# Make the appropriate folders
mkdir -p $FABRIC_CA_CLIENT_HOME/msp/admincerts
mkdir -p $FABRIC_CA_CLIENT_HOME/msp/cacerts
mkdir -p $FABRIC_CA_CLIENT_HOME/msp/keystore

# Copy the required CA certs for MSP
cp ~/fabric/CA/server/ca-cert.pem $FABRIC_CA_CLIENT_HOME/msp/cacerts

# Copy the required Admincerts for MSP
cp ~/fabric/CA/clients/$ORGANIZATION_NAME/admin/msp/signcerts/* $FABRIC_CA_CLIENT_HOME/msp/admincerts 
