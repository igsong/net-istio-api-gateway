# Istio API Gateway for Knative

This project aims to provide a way to map URL paths defined by users to Knative Services, when using Istio ingress.

Basically, combination of the conventional API gateway and Istio ingress is one way to achieve this purpose, 
but it requires one more network hop.
Therefore, this project provides more concise way to directly deliver a traffic coming with a URL path defined by users to 
Knative Services over Istio.

Currently, this is on developing.