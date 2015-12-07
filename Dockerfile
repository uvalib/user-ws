FROM centos
COPY bin/user-ws .
EXPOSE 8080
CMD ./user-ws
