# go-web-analyzer
Web application for analyzing web pages

## Setup Instructions ##
### Prerequisites ###
- Go 1.23.6 or above

### Setup and Run ###
**Step 1**: Clone the Repository
```
  git clone https://github.com/sura2k/go-web-analyzer.git`
  cd go-web-analyzer
```

**Step 2**: Install Dependencies
```
  go mod tidy
```

**Step 3**: Set Up Configuration (Optional)
1. Copy `config.yaml` from the root of `go-web-analyzer` to a preferred external location
2. Create an environment variable named `CONFIG_PATH` and set its value to the directory containing `config.yaml`
3. Modify the properties in `config.yaml` inside `CONFIG_PATH` directory as needed

*Note: This step is optional. The application will run with default settings if `CONFIG_PATH` and `config.yaml` are not provided*

#### Configurations: ####
| Configuration | Description | Default Value |
|----------|----------|----------|
| `server.port`    | Server port    | 8080    |
| `defaults.http.timeout.seconds`    | Default HTTP Timeout    | 5    |
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

**Step 7**: Submit a URL for Analysis

(Screenshot placeholder here)

**Step 8**: View Analyzed Results

(Screenshot placeholder here)

### Run Unit Tests ###
```
  cd go-web-analyzer
  go test ./... -cover
```
View Analyzed Results

### Project Structure ###
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

### System Design ###

### Challenges & Solutions ###
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

4. Configuration Management
   - Problem:
       - Some environment-specific settings, such as changing the server port, require configurable options
   - Solution:
       - Integrated the `viper` library, which supports both `yaml` configuration files and environment variables with precedence handling

### Further Considerations ###


### Code Coverage ###
(Screenshot placeholder here)

## Checklist ##

### 1. Project Structure ###
#### 1.1. Code Organization: ####
- Maintain a clear project structure with logical separation of concerns (e.g., using packages). `[DONE]`
- Ensure consistent naming conventions and code style. `[DONE]`

#### 1.2. Repository Management: ####
- Include a.gitignore file to prevent committing unnecessary files. `[DONE]`
- Use a clear commit history with descriptive messages. `[DONE]`

### 2. Documentation ###
#### 2.1. Readme: ####
- Should include a project overview, setup instructions, and usage examples. `[DONE]`
- Provide information on key components and their roles. `[TODO]`
- Include guidance on how to run the project (including prerequisites, commands, and examples). `[DONE]`
- Mention any external dependencies and how to install them. `[TODO]`
- Include details on how to run tests and what the expected outcomes are. `[TODO]`
- Discuss possible improvements or extensions to the project. `[TODO]`

#### 2.2. Additional Documentation: ####
- Create inline comments for complex code sections. `[DONE]`

### 3. Quality and Testing ###
#### 3.1. Unit Tests: ####
- Cover all major functions and methods with appropriate unit tests. `[DONE]`
- Ensure test cases include edge cases and error handling scenarios. `[DONE]`
- Use assertions to validate expected outcomes. `[WIP]`

#### 3.2. Code Coverage: ####
- Aim for a reasonable level of test coverage (e.g., >70%). `[DONE - > Covered 96% in analyzers]`
- In case of limited time try to cover at least happy path `[DONE]`

### 3.3. Code Quality Tools: ###
- Use linters (e.g., golint, go vet) to maintain consistent code quality. `[DONE using stylecheck]`
- Run static analysis to detect potential issues or vulnerabilities. `[DONE using stylecheck]`

### 4. Deployment and Operations ###
#### 4.1. Single Command for Running the Project: ####
- Provide a simple command or script to build and run the project. `[SKIPPED]`

#### 4.2. Dockerization: ####
- Only if you have knowledge with Docker `[SKIP]`
- Include a Dockerfile for containerization. `[SKIP]`
- Optionally, provide a docker-compose.yml file for complex setups. `[SKIPPED]`

#### 4.3. CI/CD Flow: ####
- Only if you have knowledge with CI CD `[SKIP]`
- Implement basic CI/CD processes (e.g., GitHub Actions, CircleCI). `[SKIPPED]`
- Set up automated builds, tests, and deployments. `[SKIP]`

#### 4.4. Environment Configuration: ####
- Use environment variables or configuration files to manage settings. `[DONE]`

#### 4.5. Logging and Monitoring: ####
- Implement basic logging to track application behavior and errors. `[DONE]`
- Optionally, integrate monitoring tools for runtime insights. `[SKIPPED]`

### 5. Other Considerations ###
#### 5.1. Error Handling and Robustness: ####
- Handle errors gracefully, providing useful error messages. `[DONE]`
- Ensure the application doesn't crash unexpectedly due to unhandled errors. `[DONE]`

#### 5.2. Code Performance and Optimization: ####
- Identify potential performance bottlenecks and optimize where necessary. `[DONE]`
- Include profiling tools for performance analysis (optional) if you have experience only. `[SKIPPED]`
- Try to utilize go routines if it makes sense. `[DONE]`
