FROM scratch

COPY rme /rme

EXPOSE 9101

ENTRYPOINT ["/rme"]
