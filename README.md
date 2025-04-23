# Gator
**A local-first RSS blog aggregator written in Go.**  

Gator is a multi-user command-line application for collecting and browsing RSS feeds. It's designed for local use and features a PostgreSQL backend, user management, and an automated feed fetcher.

---

## Requirements

Before using Gator, make sure you have:

- [Go](https://golang.org/dl/) (v1.20+)
- [PostgreSQL](https://www.postgresql.org/download/) (running locally)

---

## 🛠 Installation

### 1. Clone the Repo
```bash
git clone https://github.com/your-username/gator.git
cd gator
```
### 2. Install the CLI
```go install```
This builds and installs the gator binary into your ```$GOPATH/bin```. You should be able to run it like any CLI tool:
``` gator login your_name```

### 3. Config Setup
Gator stores configuration in a JSON file at ~/.gatorconfig.json.
Manually create this file with the following structure:
```json
{
    "db_url": "postgres://your_user:your_password@localhost:5432/gator?sslmode=disable"
}
The ```current_user``` field is automatically set by the ```login``` or ```register``` commands.

### 4. Common Commands
#### User Commands
```bash
gator register your_name    # Create a new user
gator login your_name       # Set the current user
gator users                 # List all users
```
#### Feed Commands
```bash
gator addfeed "Feed Name" https://example.com/rss
gator feeds
gator follow https://example.com/rss
gator following
gator unfollow https://example.com/rss
```
#### Aggregation
```bash
gator agg 1m                # Collect feeds every 1 minute
```
This continuously fetches and stores new posts from feeds you're following.
#### Browsing Posts
```bash
gator browse [limit]        # View recent posts (default limit: 2)
```
#### Reset
```bash
gator reset
```
This will delete all users, feeds, and posts.

### PostgresSQL Setup (Quick Ref)
To create the required ```gator``` database:
```bash
createdb gator
```
To check your tables:
```bash
psql gator
\dt
```