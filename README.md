# Backend Assignment

This repository contains the source code for the backend assignment.

## Installation

### Requirements

- Go programming language
- Docker (optional, if you prefer to run PostgreSQL in a container)

### Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/emirhanusta/backend-assesment.git
    ```

2. Navigate to the project directory:

    ```bash
    cd backend-assignment
    ```

3. **If you prefer to use Docker for PostgreSQL:**

    - Navigate to the file directory:
      
    ```bash
    cd scripts
    ```

    - Then, run the following command in Git Bash to start the PostgreSQL database in Docker
  
   ```bash
    sh backend_assesment_db.sh
    ```

5. **If you prefer to use a locally installed PostgreSQL:**

    - Ensure you have PostgreSQL installed and running on your local machine.
  
    - Create a PostgreSQL database named `report`.

    - Run the SQL script `assesment.sql` located in the `scripts` directory to set up the required tables and insert sample data.

6. Make sure you have the following PostgreSQL connection configuration:

    ```yaml
    Host: localhost
    Port: 5432
    Database: report
    Username: postgres
    Password: postgres
    ```

7. Run the backend server:

    The server will start running on http://localhost:8080.

    ```bash
    go run main.go
    ```

## Usage

Access the backend server through the provided API endpoints. The server provides endpoints for querying reports with pagination.

### API Endpoints

- **POST http://localhost:8080/assignment/query?page=1&page_size=5**: Query reports with pagination. Send a JSON payload with filtering and ordering criteria.

## Sample Request Body

```json
{
    "filters": [
        {
            "column": "main_symbol",
            "value": ["LOC407835"]
        },
        {
            "column": "main_dp",
            "value": 95
        }
    ],
    "ordering": [
        {
            "column": "main_af_vcf",
            "direction": "DESC"
        }
    ]
}
```
- The above request body demonstrates how to filter reports by the main_symbol column with a value of "LOC407835" , main_dp column with a value of 95  and order the results by main_af_vcf in descending order

## Sample Response
```json
{
    "page": 1,
    "page_size": 10,
    "count": 3,
    "results": [
        {
            "row": 72,
            "main_uploaded_variation": "7_129126717_AAGACGATGACTTCG/-",
            "main_existing_variation": "rs759583159",
            "main_symbol": "LOC407835",
            "main_af_vcf": 0.4835,
            "main_dp": 95,
            "details2_provean": "",
            "links_mondo": "",
            "links_pheno_pubmed": ""
        },
        {
            "row": 69,
            "main_uploaded_variation": "7_129126742_AGA/-",
            "main_existing_variation": "rs780606902",
            "main_symbol": "LOC407835",
            "main_af_vcf": 0.4632,
            "main_dp": 95,
            "details2_provean": "",
            "links_mondo": "",
            "links_pheno_pubmed": ""
        },
        {
            "row": 70,
            "main_uploaded_variation": "7_129126740_T/C",
            "main_existing_variation": "rs746112504",
            "main_symbol": "LOC407835",
            "main_af_vcf": 0.4632,
            "main_dp": 95,
            "details2_provean": "",
            "links_mondo": "",
            "links_pheno_pubmed": ""
        }
    ]
}
```
- The response contains paginated results with the specified page size and the total count of reports. Each result object includes detailed information about a report


