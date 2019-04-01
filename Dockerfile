ARG SYSL2_VERSION

FROM anzbank/sysl2:${SYSL2_VERSION}

ENV SYSLGEN_BASE=/syslgen-examples
COPY . .
