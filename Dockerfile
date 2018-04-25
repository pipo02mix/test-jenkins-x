FROM scratch
EXPOSE 8080
ENTRYPOINT ["/test-jenkins-x"]
COPY ./bin/ /