# FACEIT DevOps Challenge

You have been asked to create the infrastructure for running a web application on a cloud platform of your preference (Google Cloud Platform preferred, AWS or Azure are also fine).

The [web application](test-app/README.md) can be found in the `test-app/` directory. Its only requirements are to be able to connect to a PostgreSQL database and perform PING requests.    

The goal of the challenge is to demonstrate hosting, managing, documenting and scaling a production-ready system.

This is not about website content or UI.

## Requirements

- Deliver the tooling to set up the application using Terraform on the cloud platform of your choice (free tiers are fine)
- Provide basic architecture diagrams of your solution and documentation on how to initialise the infrastructure
- Document a mechanism for scaling the service, and ensuring high availability
- Describe a possible solution for CI and/or CI/CD in order to release a new version of the application to production without any downtime
- Explain your choices

## Evaluation

We are looking for:

- _your_ solution to this problem, without constraints on platforms/technologies used _except for Terraform_.
- _your_ design, and reasoning on the solution, more so than implementation details.

We are not looking:

- specific platforms.
- shiny new technologies.

## General guidance

- We recommend using this repository as a starting point, you can clone it and add your code/docs to that repository
- Please do no open pull request with your challenge against **this repository**
- Submission of the challenge can be done either via your own public repository or zip file containing `.git` folder
