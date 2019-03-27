ARG SYSL2_VERSION

FROM anzbank/sysl2:${SYSL2_VERSION}

ENV SYSLGEN_BASE=/syslgen-examples
COPY grammars $SYSLGEN_BASE/grammars
COPY transforms $SYSLGEN_BASE/transforms
