# go-web-analyzer
Web application for analyzing web pages

`[TODO]`
`REDME instructions and other details to be added`


# Checklist #

## 1. Project Structure ##
### 1.1. Code Organization: ###
- Maintain a clear project structure with logical separation of concerns (e.g., using packages). `[DONE]`
- Ensure consistent naming conventions and code style. `[DONE]`

### 1.2. Repository Management: ###
- Include a.gitignore file to prevent committing unnecessary files. `[DONE]`
- Use a clear commit history with descriptive messages. `[DONE]`

## 2. Documentation ##
### 2.1. Readme: ###
- Should include a project overview, setup instructions, and usage examples. `[TODO]`
- Provide information on key components and their roles. `[TODO]`
- Include guidance on how to run the project (including prerequisites, commands, and examples). `[TODO]`
- Mention any external dependencies and how to install them. `[TODO]`
- Include details on how to run tests and what the expected outcomes are. `[TODO]`
- Discuss possible improvements or extensions to the project. `[TODO]`

### 2.2. Additional Documentation: ###
- Create inline comments for complex code sections. `[DONE]`

## 3. Quality and Testing ##
### 3.1. Unit Tests: ###
- Cover all major functions and methods with appropriate unit tests. `[WIP]`
- Ensure test cases include edge cases and error handling scenarios. `[WIP]`
- Use assertions to validate expected outcomes. `[WIP]`

### 3.2. Code Coverage: ###
- Aim for a reasonable level of test coverage (e.g., >70%). `[WIP - > Covered 75%+ in analyzers]`
- In case of limited time try to cover at least happy path `[WIP]`

### 3.3. Code Quality Tools: ###
- Use linters (e.g., golint, go vet) to maintain consistent code quality. `[DONE using stylecheck]`
- Run static analysis to detect potential issues or vulnerabilities. `[DONE using stylecheck]`

## 4. Deployment and Operations ##
### 4.1. Single Command for Running the Project: ###
- Provide a simple command or script to build and run the project. `[TODO]`

### 4.2. Dockerization: ###
- Only if you have knowledge with Docker `[SKIP]`
- Include a Dockerfile for containerization. `[SKIP]`
- Optionally, provide a docker-compose.yml file for complex setups. `[SKIP]`

### 4.3. CI/CD Flow: ###
- Only if you have knowledge with CI CD `[SKIP]`
- Implement basic CI/CD processes (e.g., GitHub Actions, CircleCI). `[SKIP]`
- Set up automated builds, tests, and deployments. `[SKIP]`

### 4.4. Environment Configuration: ###
- Use environment variables or configuration files to manage settings. `[DONE]`

### 4.5. Logging and Monitoring: ###
- Implement basic logging to track application behavior and errors. `[DONE]`
- Optionally, integrate monitoring tools for runtime insights. `[SKIP]`

## 5. Other Considerations ##
### 5.1. Error Handling and Robustness: ###
- Handle errors gracefully, providing useful error messages. `[DONE]`
- Ensure the application doesn't crash unexpectedly due to unhandled errors. `[DONE]`

### 5.2. Code Performance and Optimization: ###
- Identify potential performance bottlenecks and optimize where necessary. `[TODO]`
- Include profiling tools for performance analysis (optional) if you have experience only. `[SKIP]`
- Try to utilize go routines if it makes sense. `[DONE]`
