#!/bin/sh

# mkdir ~/apiman-2.0.0.Final
# cd ~/apiman-2.0.0.Final
# curl -L https://download.jboss.org/wildfly/20.0.1.Final/wildfly-20.0.1.Final.zip -o wildfly-20.0.1.Final.zip
# curl -L https://github.com/apiman/apiman/releases/download/2.0.0.Final/apiman-distro-wildfly-2.0.0.Final-overlay.zip -o apiman-distro-wildfly-2.0.0.Final-overlay.zip
# unzip wildfly-20.0.1.Final.zip
# unzip -o apiman-distro-wildfly-2.0.0.Final-overlay.zip -d wildfly-20.0.1.Final
# cd wildfly-20.0.1.Final
# ./bin/standalone.sh -c standalone-apiman.xml

docker pull apiman/on-wildfly:2.0.0.Final
docker run -it -p 8080:8080 -p 8443:8443 apiman/on-wildfly:2.0.0.Final