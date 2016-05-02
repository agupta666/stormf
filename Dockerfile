FROM       scratch
MAINTAINER Anirban Gupta <agupta666@gmail.com>
ADD        stormf stormf
EXPOSE     3000
ENTRYPOINT ["/stormf"]
