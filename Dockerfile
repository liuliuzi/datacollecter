FROM busybox
WORKDIR /app
ADD bin/datacollecter /app/
EXPOSE 8060
CMD ["./datacollecter"]