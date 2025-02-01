 # MVP Features for FluxAPI

 1. API Request Execution
    - HTTP Methods: Support for GET, POST, PUT, DELETE, and PATCH requests
    - Request Headers: Custom header management
    - Request parameters: Support for query parameters and body data (JSON, form data, etc.)
    - Environment variables: Easy environment setup (e.g, dev, production) to switch between different API endpoints
2. Test Generation from Code
    - Auto-generate tests:
        - Parse OpenAPI specification and auto-generate tests based on the routes, defined request/response structures, and any property validation rules
        - Parse given code base to generate tests from source code
            - Not sure if this would be in MVP...
    - Assertion Support
        - Automatically add common assertions (status codes, response structure, content type, etc)
        - CLI command: ```fluxapi generate-tests```
            - Generates a basic test suite for the API
        - *Notes*: It will be hard to have context and business rules without code scanning and AI... is this a premium feature for later on?
3. Seed Data Generation and Execution
    - Database seeding:
        - Automatically generate and execute seed data scripts to populate a database with realistic data for testing
    - Custom scripts:
        - Ability to write custom seed data scripts for complex test setups
    - CLI Command: ```fluxapi seed --env=staging```
4. Basic Automation Suite
    - Run tests in sequence:
        - Ability to execute multiple API tests in sequence, useful for regression testing
    - Simple Scheduling:
        - Run tests periodically (e.g., ```fluxapi run --schedule every-24hrs```)
    - CI/CD Integration: Basic integration into CI pipelines (e.g., GH actions or ADO pipelines)
    - CLI Command: ```fluxapi run --suite=test-suite.json```
5. Basic Load Testing:
    - Simulate load:
        - Allow users to simulate a defined number of requests per second to test API performance under stress
    - Response Time Tracking:
        - Track and log API response times during load tests
    - CLI Command: ```fluxapi load-test --rps=1000 --url=https://example.com/api```
6. CLI-Based Reporting
    - Test Results:
        - After running tests, display pass/fail results, response times, and basic error reports in the terminal
    - Basic logs:
        - View the raw logs for requests/responses and identify any failures
    - CLI Command ```fluxapi report```

## Next Steps:
1. Set up CLI interface
    - Core command-line interface
    - Focus on core commands:
        - test
        - generate-tests
        - seed
        - load-test
2. Build Basic Test Engine
    - Define the structure for storing tests (local files? JSON or YAML?)
        - I think we start with this for the CLI
        - The pro version with UI will store them in the cloud in a more performant format
            - This will also allow the cloud synchronization for teams
            - A user can then upload their local files to be converted into the cloud versions if they wish?
    - Auto generate tests
    - Executing requests
3. Initial Reporting system
    - Output basic pass/fail reports in the terminal
4. Performance Testing Integration
    - Include basic load testing with response time tracking
5. CI/CD Integration
    - Ensure that FluxAPI can be easily integrated into CI pipelines for automated testing

## Tech Stack:
- Golang for the backend
- Use the Cobra package for golang to help build the CLI
- UI: Next.js + tailwindcss


## Future Ideas:
- Slack integration
    - Send alerts when a test fails or when load testing exceeds thresholds
- Performance graphs
    - Once the UI is ready, add visual performance graphs for load testing results
- Advanced scheduling
    - Set up test suites to run at specific times or intervals
- Test History and Analytics
    - Track the history of API tests over time (to be included in the UI version)
            
