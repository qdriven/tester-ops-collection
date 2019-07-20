FROM 3.7.4-alpine3.9

LABEL maintainer="patrick"

RUN apk add --no-cache ansible openssh fping curl rsyslog iputils-ping netbase net-tools

WORKDIR /root/data

RUN mkdir -p /root/.ssh && chmod 700 /root/.ssh && mkdir /etc/ansible

COPY common/ansible.cfg /etc/ansible/ansible.cfg

COPY ["common/id_rsa", "common/id_rsa.pub", "/root/.ssh/"]
COPY ["common/.vimrc", "common/.profile", "/root/"]

RUN chmod 644 /root/.ssh/id_rsa.pub && chmod 600 /root/.ssh/id_rsa


