# s2i-do288-httpd
FROM registry.access.redhat.com/ubi8/ubi:8.4

# TODO: Put the maintainer name in the image metadata
# LABEL maintainer="Your Name <your@email.com>"

LABEL Component="httpd" \
      Name="s2i-do288-httpd" \
      Version="1.0" \
      Release="1"

# TODO: Rename the builder environment variable to inform users about application you provide them
# ENV BUILDER_VERSION 1.0

# TODO: Set labels used in OpenShift to describe the builder image
#LABEL io.k8s.description="Platform for building xyz" \
#      io.k8s.display-name="builder x.y.z" \
#      io.openshift.expose-services="8080:http" \
#      io.openshift.tags="builder,x.y.z,etc."

# Labels consumed by RHOCP
LABEL io.k8s.description="A basic Apache HTTP Server S2I builder image" \
      io.k8s.display-name="Apache HTTP Server S2I builder image for DO288" \
      io.openshift.expose-services="8080:http" \
      io.openshift.s2i.scripts-url="image:///usr/libexec/s2i" \
      io.openshift.tags="builder, httpd, httpd24"
    

# TODO: Install required packages here:
# RUN yum install -y ... && yum clean all -y

ENV DOCROOT /var/www/html

RUN   yum install -y --nodocs --disableplugin=subscription-manager httpd && \ 
      yum clean all --disableplugin=subscription-manager -y && \
      echo "This is the default index page from the s2i-do288-httpd S2I builder image." > ${DOCROOT}/index.html 

RUN sed -i "s/Listen 80/Listen 8080/g" /etc/httpd/conf/httpd.conf

# TODO (optional): Copy the builder files into /opt/app-root
# COPY ./<builder_folder>/ /opt/app-root/

# TODO: Copy the S2I scripts to /usr/libexec/s2i, since openshift/base-centos7 image
# sets io.openshift.s2i.scripts-url label that way, or update that label
COPY ./s2i/bin/ /usr/libexec/s2i

# TODO: Drop the root user and make the content of /opt/app-root owned by user 1001
# RUN chown -R 1001:1001 /opt/app-root

# This default user is created in the openshift/base-centos7 image
USER 1001

# TODO: Set the default port for applications built using this image
EXPOSE 8080

# TODO: Set the default CMD for the image
CMD ["/usr/libexec/s2i/usage"]
