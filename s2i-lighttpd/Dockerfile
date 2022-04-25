# lighttpd-centos7
FROM openshift/base-centos7

# TODO: Put the maintainer name in the image metadata
# LABEL maintainer="Your Name <your@email.com>"
LABEL maintainer="El Shebi <juansebasrodriguez2002@gmail.com>"
# TODO: Rename the builder environment variable to inform users about application you provide them
# ENV BUILDER_VERSION 1.0
ENV LIGHTTPD_VERSION=1.4.36
# TODO: Set labels used in OpenShift to describe the builder image
LABEL io.k8s.description="Platform for building static HTML files" \
      io.k8s.display-name="Lighthttp 1.4.36" \
      io.openshift.expose-services="8080:http" \
      io.openshift.tags="builder,html,lighthttp"

# TODO: Install required packages here:
# RUN yum install -y ... && yum clean all -y
# Install the required software, namely Lighttpd and
RUN yum -y install epel-release
RUN yum install -y lighttpd && yum clean all -y
# clean yum cache files, as they are not needed and will only make the image bigger in the end


#RUN gem install asdf

# TODO (optional): Copy the builder files into /opt/app-root
# COPY ./<builder_folder>/ /opt/app-root/

# TODO: Copy the S2I scripts to /usr/libexec/s2i, since openshift/base-centos7 image
# sets io.openshift.s2i.scripts-url label that way, or update that label
COPY ./s2i/bin/ /usr/libexec/s2i

# TODO: Drop the root user and make the content of /opt/app-root owned by user 1001
# RUN chown -R 1001:1001 /opt/app-root

# Copy the lighttpd configuration file
COPY ./etc/ /opt/app-root/etc

RUN chown -R 1001:1001 /opt/app-root


# This default user is created in the openshift/base-centos7 image
USER 1001

# TODO: Set the default port for applications built using this image
EXPOSE 8080

# TODO: Set the default CMD for the image
#docker run
CMD ["/usr/libexec/s2i/usage"]
