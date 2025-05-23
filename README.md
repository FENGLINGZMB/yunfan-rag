# Yunfan RAG

This project is a Retrieval Augmented Generation (RAG) system for Yunfan, utilizing Milvus for vector storage and PostgreSQL for metadata.

## Getting Started

This section will guide you through setting up and running the Yunfan RAG project.

### Prerequisites

- Python 3.8+
- Docker
- Docker Compose

### Setup

1.  **Clone the repository:**
    ```bash
    git clone <repository_url>
    cd <repository_name>
    ```

2.  **Configure Milvus:**
    - Milvus is used as the vector database.
    - Instructions for setting up Milvus using Docker Compose are typically provided in a `docker-compose.yml` file within the project.
    - Ensure Milvus is running before starting the application.
    - Default connection parameters for Milvus are usually:
        - Host: `localhost`
        - Port: `19530`
    - Update configuration files if your Milvus instance is running elsewhere.

3.  **Configure PostgreSQL:**
    - PostgreSQL is used for storing metadata.
    - Instructions for setting up PostgreSQL using Docker Compose are often included in the `docker-compose.yml` file.
    - Ensure PostgreSQL is running.
    - Default connection parameters for PostgreSQL are typically:
        - Host: `localhost`
        - Port: `5432`
        - User: `user`
        - Password: `password`
        - Database: `yunfan_rag_db`
    - Update configuration files (e.g., `.env` or `config.py`) with your PostgreSQL credentials and database name. Create the database if it doesn't exist.

4.  **Install dependencies:**
    ```bash
    pip install -r requirements.txt
    ```

### Running the Application

1.  **Set up environment variables:**
    - Create a `.env` file if one is not already present, and populate it with necessary configurations (e.g., API keys, database URIs). Refer to a `.env.example` file if available.

2.  **Run database migrations (if applicable):**
    - If the project uses a migration tool like Alembic or Django migrations:
      ```bash
      # Example for Alembic
      alembic upgrade head
      ```

3.  **Start the application:**
    ```bash
    python main.py
    ```
    (Or the relevant command to start your application server)

## Installation

Instructions on how to install and set up the project.

## Usage

This section describes how to interact with the Yunfan RAG system once it is running.

### Querying the System

The primary way to use the RAG system is by sending queries to its API endpoint (assuming it exposes one).

**Example: Querying via API (Conceptual)**

If the system has an API endpoint, you might interact with it using a tool like `curl` or a client library in your preferred programming language.

```bash
# Example using curl
curl -X POST http://localhost:8000/api/query \
     -H "Content-Type: application/json" \
     -d '{
           "query": "What are the latest advancements in AI?",
           "top_k": 5
         }'
```

**Expected Response:**

The system would return a JSON response containing the retrieved documents and the generated answer.

```json
{
  "query": "What are the latest advancements in AI?",
  "generated_answer": "Recent advancements in AI include improvements in large language models, generative adversarial networks, and reinforcement learning techniques, leading to more capable and versatile AI systems.",
  "retrieved_documents": [
    {
      "id": "doc_123",
      "source": "research_paper_XYZ.pdf",
      "text_snippet": "Large language models have shown remarkable progress in natural language understanding and generation...",
      "score": 0.92
    },
    {
      "id": "doc_456",
      "source": "blog_post_ABC.com",
      "text_snippet": "Generative adversarial networks are being used to create realistic images, videos, and even music...",
      "score": 0.88
    }
    // ... other documents
  ]
}
```

### Using a Client Application (If available)

If a client application (e.g., a web interface or a command-line tool) is provided, refer to its specific documentation for usage instructions.

**Example: Command-Line Interface (Conceptual)**

```bash
python cli_client.py --query "Tell me about Milvus performance."
```

This would then display the answer and relevant sources in the console.

### Key Parameters for Querying

When querying the system, you might be able to specify parameters like:

-   `query` (string): The question or topic you want information about.
-   `top_k` (integer, optional): The number of relevant documents to retrieve from Milvus. Defaults to a system-defined value (e.g., 3 or 5).
-   `filters` (object, optional): Metadata filters to apply to the search (e.g., filter by document source, date range, etc.). The exact structure will depend on the system's implementation.

Refer to the API documentation or `main.py --help` (if applicable) for detailed information on available endpoints and parameters.

## Contributing

We welcome contributions to the Yunfan RAG project! Whether it's bug fixes, new features, or improvements to documentation, your help is appreciated.

### How to Contribute

1.  **Fork the Repository:**
    Start by forking the main repository to your GitHub account.

2.  **Clone Your Fork:**
    Clone your forked repository to your local machine:
    ```bash
    git clone https://github.com/YOUR_USERNAME/yunfan-rag.git
    cd yunfan-rag
    ```

3.  **Create a New Branch:**
    Create a new branch for your changes. Choose a descriptive branch name (e.g., `feature/add-new-embedding-model` or `fix/query-api-bug`).
    ```bash
    git checkout -b your-branch-name
    ```

4.  **Make Your Changes:**
    Implement your feature, fix the bug, or improve the documentation.
    - Ensure your code follows the project's coding style (if one is defined, e.g., via a linter like Flake8 or Black).
    - Write clear and concise commit messages.

5.  **Add/Update Tests:**
    - If you're adding a new feature, please include unit tests or integration tests that cover its functionality.
    - If you're fixing a bug, add a test that reproduces the bug and verifies your fix.
    - Ensure all tests pass before submitting your contribution.
    ```bash
    # Example: pytest tests/
    ```

6.  **Update Documentation (if applicable):**
    - If your changes affect how the system is used or configured, please update the `README.md` or other relevant documentation files.

7.  **Commit Your Changes:**
    ```bash
    git add .
    git commit -m "Your descriptive commit message"
    ```

8.  **Push to Your Fork:**
    Push your changes to your forked repository on GitHub:
    ```bash
    git push origin your-branch-name
    ```

9.  **Submit a Pull Request (PR):**
    - Go to the original Yunfan RAG repository on GitHub.
    - You should see a prompt to create a Pull Request from your recently pushed branch.
    - Click "Compare & pull request".
    - Provide a clear title and description for your PR, explaining the changes you've made and why.
    - Link any relevant issues (e.g., "Closes #123").
    - Submit the PR.

### Contribution Guidelines

-   **Code Style:** Follow existing code style. If a linter configuration (e.g., `.flake8`, `pyproject.toml` for Black) is present, adhere to it.
-   **Testing:** Ensure your changes don't break existing tests and add new tests for new functionality.
-   **Communication:** If you plan to work on a large feature, it's a good idea to open an issue first to discuss your approach.
-   **Be Respectful:** Treat other contributors with respect.

Thank you for contributing!

## License

Information about the project's license.
