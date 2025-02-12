# go-web-analyzer
Web application for analyzing web pages

# Table of Contents #
     
1. [Design Overview](#1-design-overview)

     1.1. [Project Structure](#11-project-structure)

     1.2. [Key Components](#12-key-components)

     1.3. [System Design](#13-system-design)

     1.4. [Challenges & Solutions (Already Adressed)](#14-challenges--solutions-already-adressed)

     1.5. [Further Improvements (Suggestions)](#15-further-improvements-suggestions)

     1.6. [Concerns and Assumptions](#16-concerns-and-assumptions)
   
2. [Setup the Project](#2-setup-the-project)
   
     2.1. [Prerequisites](#21-prerequisites)
     
     2.2. [Setup and Run](#22-setup-and-run)
     
     2.2.3. [Configurations](#configurations)
     
     2.3. [Changing Configurations](#23-changing-configurations)
     
     2.4. [Run Unit Tests](#24-run-unit-tests)
     
     2.5. [Code Coverage](#25-code-coverage)
   
4. [Checklist](#3-checklist)


# 1. Design Overview #
### 1.1. Project Structure ###
```
/
  ├── config/                 - Contains structs and functions for managing application configuration 
  ├── controllers/            - Contains all controller-related functions           
  │    └── rest/              - Contains REST controllers (Note: No REST controllers at the moment) 
  │    └── view/              - Contains view controllers for rendering UI  
  ├── services/               - Handles business logic and core functionalities
  │    └── analyzers/         - Implements analyzers and related functions 
  │    └── utils/             - Provides shared utility and helper functions 
  ├── templates/              - Stores HTML view templates 
  │
  └── main.go                 - Entry point of the application 
  └── config.yaml             - A reference-only configuration file
```

### 1.2. Key Components ###
1. `/main.go` - Application entry point
2. `/controller/view/view_controller.go` - View Controller handles page's GET and submit POST requests
3. `/services/analyzers/analyzer_manager.go` - Act as the analyzer workflow orchestrator
4. `/services/analyzers/analyzer_preprocessor.go` - A Preprocessor which validates input, fetch html content and transform, etc. Invoked by the analyzer_manager
5. `/services/analyzers/analyzer_executor.go` - Executor all the analyzers parallaly. Invoked by the analyzer_manager
6. `/services/analyzers/*_analyzer.go` - All the available analyzer implementations

### 1.3. System Design ###
![LucyTech_Analyzer drawio](https://github.com/user-attachments/assets/bebb87b6-1af5-4695-b7f9-3cb201d24f4b)

![Untitled](https://github.com/user-attachments/assets/e95b43a4-84b4-41c4-a3c8-3c64336a4318)


### 1.4. Challenges & Solutions (Already Adressed) ###
*1. Analyzer execution*
   - Problem:
       - When analyzers run sequentially, the process takes significantly longer as it is a blocking operation
   - Solution
       - Run all analyzers in parallel using goroutines
   
*2. Improved Link Accessibility Checks in `LinksAnalyzer`*
   - Problem:
       - When `LinksAnalyzer` checks the accessibility of multiple `<a href="...">` links in the document, execution time increases as the number of links grows
   - Solution 1:
       - Execute link accessibility checks in parallel using goroutines
   - Problem with Solution 1:
       - If too many links are processed simultaneously, high CPU/RAM usage may lead to system crashes
   - Solution 2:
       - Process links in batches. Each batch runs in parallel, and the next batch starts after the previous one completes
       - The default batch size is 10 and can be adjusted using the configuration property `analyzers.LinksAnalyzer.link-health-check.batch_size`
   
*3. Enhanced HTML Content Fetching from URLs*
   - Problem:
       - Go's default `http.Get()` fetches only the initial HTML content, often **before JavaScript execution** or **before a proxy processes the request**. As a result, the fetched HTML may not match what is displayed in a browser.
   - Solution:
       - Use `chromedp`, a JavaScript-enabled headless browser library, to fetch the complete HTML after JavaScript has fully loaded

*4. Configuration Management*
   - Problem:
       - Some environment-specific settings, such as changing the server port, require configurable options
   - Solution:
       - Integrated the `viper` library, which supports both `yaml` configuration files and environment variables with precedence handling

### 1.5. Further Improvements (Suggestions) ###
1. Move the frontend to a separate project, possibly using React or Angular
2. Replace view controllers with REST controllers
3. Add rate limiters to control the number of requests
4. If the number of analyzers grows, process them in batches when a set limit is exceeded
5. Allow the backend to enable or disable specific analyzers through configuration
6. Update the UI and APIs to let users choose which analyzers to run

### 1.6. Concerns and Assumptions  ###
1. Hidden heading tags (`<h1>`, `<h2>`, etc.) and link tags (`<a>`) are treated as valid since handling all edge cases is complex
2. Some login forms, especially Ajax-based logins, may not use standard `<form>` tags, so they are not detected
3. Hidden forms (`<form>`) with proper login fields are still considered valid login forms
4. `vipe` library gives higher priority to environment variables over configuration files by default


# 2. Setup the Project #
### 2.1 Prerequisites ###
- Go 1.23.6 or above

### 2.2. Setup and Run ###
**Step 1**: Clone the Repository
```
  git clone https://github.com/sura2k/go-web-analyzer.git`
  cd go-web-analyzer
```

**Step 2**: Install Dependencies

Run the following command to install necessary dependencies:
```
  go mod tidy
```

Note: The following external dependencies will be installed:
  - `github.com/chromedp/chromedp v0.12.1` – Headless browser for fetching fully rendered HTML
  - `github.com/spf13/viper v1.19.0` – Configuration management with support for `yaml`, `toml`, etc and `environment variables`
  - `github.com/stretchr/testify v1.10.0` – Testing utilities for assertions

**Step 3**: Setup Configuration (Optional)
1. Copy `config.yaml` from the root of `go-web-analyzer` to a preferred external location
2. Create an environment variable named `CONFIG_PATH` and set its value to the directory containing `config.yaml`
3. Modify the properties in `config.yaml` inside `CONFIG_PATH` directory as needed

*Note: This step is optional. The application will run with default settings if `CONFIG_PATH` and `config.yaml` are not provided*

#### Configurations: ####
| Configuration | Description | Default Value |
|----------|----------|----------|
| `server.port`    | Server port    | 8080    |
| `defaults.http.timeout.seconds`    | Default HTTP Timeout (in seconds)   | 5    |
| `analyzers.LinksAnalyzer.link-health-check.batch_size` | Link Helath Check Batch Size    |  10    |
  
**Step 4**: Run from Source
```
  cd go-web-analyzer
  go run main.go
```

**Step 5**: Verify via Logs
If the application starts successfully, you should see the following logs in the console:
```
2025/02/11 16:34:06 Config file loaded successfully
2025/02/11 16:34:06 Configurations loaded
2025/02/11 16:34:06 Server started on port  8080
```
*Note: If you need to use a different port, update the `server.port` property in `config.yaml`*

**Step 6**: Access the Application

Open your browser and navigate to: `http://localhost:8080`

![image](https://github.com/user-attachments/assets/61830fbb-4aa7-43af-9efd-176b671ab972)

**Step 7**: Submit a URL for Analysis

![image](https://github.com/user-attachments/assets/ff5ab3f3-c115-4b91-ab66-f438a12c208d)

You will see logs in your console as in shown below:
![image](https://github.com/user-attachments/assets/59c3c51e-b8b4-4e8f-baa4-992f3fb0a912)

**Step 8**: View Analyzed Results

![image](https://github.com/user-attachments/assets/68209fd3-e6cb-4cf6-b144-feab38d082e4)

*Note: If you notice or encountered that some of the accessible links are returned as not-accessible, try to increase the http time out via `defaults.http.timeout.seconds` configuration property.*

### 2.3. Changing Configurations ###

If you need to change certain configuration properties, make sure to update them in `config.yaml` located in `CONFIG_PATH` directory and restart the application

### 2.4. Run Unit Tests ###
```
  cd go-web-analyzer
  go test ./... -cover
```

### 2.5. Code Coverage ###

![image](https://github.com/user-attachments/assets/1e203aab-f635-48d1-8fe7-8029da840215)

-----------------------------------------------------------------------

# 3. Checklist #

### 3.1. Project Structure ###
#### 3.1.1. Code Organization: ####
- Maintain a clear project structure with logical separation of concerns (e.g., using packages). `[DONE]`
- Ensure consistent naming conventions and code style. `[DONE]`

#### 3.1.2. Repository Management: ####
- Include a.gitignore file to prevent committing unnecessary files. `[DONE]`
- Use a clear commit history with descriptive messages. `[DONE]`

### 3.2. Documentation ###
#### 3.2.1. Readme: ####
- Should include a project overview, setup instructions, and usage examples. `[DONE]`
- Provide information on key components and their roles. `[DONE]`
- Include guidance on how to run the project (including prerequisites, commands, and examples). `[DONE]`
- Mention any external dependencies and how to install them. `[DONE]`
- Include details on how to run tests and what the expected outcomes are. `[DONE]`
- Discuss possible improvements or extensions to the project. `[DONE]`

#### 3.2.2. Additional Documentation: ####
- Create inline comments for complex code sections. `[DONE]`

### 3.3. Quality and Testing ###
#### 3.3.1. Unit Tests: ####
- Cover all major functions and methods with appropriate unit tests. `[DONE]`
- Ensure test cases include edge cases and error handling scenarios. `[DONE]`
- Use assertions to validate expected outcomes. `[DONE]`

#### 3.3.2. Code Coverage: ####
- Aim for a reasonable level of test coverage (e.g., >70%). `[DONE - > Covered 96% in analyzers]`
- In case of limited time try to cover at least happy path `[DONE]`

### 3.3.3. Code Quality Tools: ###
- Use linters (e.g., golint, go vet) to maintain consistent code quality. `[DONE using stylecheck]`
- Run static analysis to detect potential issues or vulnerabilities. `[DONE using stylecheck]`

### 3.4. Deployment and Operations ###
#### 3.4.1. Single Command for Running the Project: ####
- Provide a simple command or script to build and run the project. `[SKIPPED]`

#### 3.4.2. Dockerization: ####
- Only if you have knowledge with Docker `[SKIPPED]`
- Include a Dockerfile for containerization. `[SKIPPED]`
- Optionally, provide a docker-compose.yml file for complex setups. `[SKIPPED]`

#### 3.4.3. CI/CD Flow: ####
- Only if you have knowledge with CI CD `[SKIPPED]`
- Implement basic CI/CD processes (e.g., GitHub Actions, CircleCI). `[SKIPPED]`
- Set up automated builds, tests, and deployments. `[SKIPPED]`

#### 3.4.4. Environment Configuration: ####
- Use environment variables or configuration files to manage settings. `[DONE]`

#### 3.4.5. Logging and Monitoring: ####
- Implement basic logging to track application behavior and errors. `[DONE]`
- Optionally, integrate monitoring tools for runtime insights. `[SKIPPED]`

### 3.5. Other Considerations ###
#### 3.5.1. Error Handling and Robustness: ####
- Handle errors gracefully, providing useful error messages. `[DONE]`
- Ensure the application doesn't crash unexpectedly due to unhandled errors. `[DONE]`

#### 3.5.2. Code Performance and Optimization: ####
- Identify potential performance bottlenecks and optimize where necessary. `[DONE]`
- Include profiling tools for performance analysis (optional) if you have experience only. `[SKIPPED]`
- Try to utilize go routines if it makes sense. `[DONE]`
