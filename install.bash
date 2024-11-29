#!/bin/bash

# Todo CLI Management Script

# Determine the script's directory
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_DIR="$HOME/dev/TODOCLI"  # Ensure the path is correctly resolved

# Color codes for enhanced output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to display usage information
usage() {
    echo -e "${YELLOW}Todo CLI Management Script${NC}"
    echo "Usage:"
    echo "  todocli build         - Compile the Go application"
    echo "  todocli install       - Install the application globally"
    echo "  todocli run           - Run the application directly"
    echo "  todocli help          - Show this help message"
    echo "  todocli uninstall     - Remove the installed application"
    echo "  todocli setup-db      - Create the database and table"
}

# Build the application
build() {
    echo -e "${GREEN}Building Todo CLI Application...${NC}"
    cd "$PROJECT_DIR" || exit 1
    go build -o todocli
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}Build successful!${NC}"
    else
        echo -e "${RED}Build failed!${NC}"
        exit 1
    fi
}

# Install the application globally
install() {
    build
    echo -e "${GREEN}Installing Todo CLI...${NC}"
    sudo mv todocli /usr/local/bin/todocli
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}Installation complete! You can now use 'todocli' command.${NC}"
    else
        echo -e "${RED}Installation failed!${NC}"
        exit 1
    fi
}

# Run the application
run() {
    cd "$PROJECT_DIR" || exit 1
    go run ./ "$@"
}

# Uninstall the application
uninstall() {
    echo -e "${YELLOW}Uninstalling Todo CLI...${NC}"
    sudo rm /usr/local/bin/todocli
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}Uninstallation complete!${NC}"
    else
        echo -e "${RED}Uninstallation failed!${NC}"
        exit 1
    fi
}

# Setup the database
setup_db() {
    echo -e "${GREEN}Setting up the database...${NC}"

    # Ensure the DBUSER and DBPASS are set
    if [ -z "$DBUSER" ] || [ -z "$DBPASS" ]; then
        echo -e "${RED}DBUSER and DBPASS environment variables are not set!${NC}"
        echo -e "Please set them using:\nexport DBUSER=<your_db_user>\nexport DBPASS=<your_db_password>"
        exit 1
    fi

    # SQL to execute
    SQL=$(cat <<EOF
DROP TABLE IF EXISTS todo;
CREATE TABLE todo (
  id           INT AUTO_INCREMENT NOT NULL,
  title        VARCHAR(255) NOT NULL,
  completed    BOOLEAN NOT NULL DEFAULT FALSE,
  created_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  completed_at DATETIME NULL,
  PRIMARY KEY (id)
);
EOF
)

    # Execute the SQL using MySQL
    echo "$SQL" | mysql -u"$DBUSER" -p"$DBPASS"
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}Database setup completed successfully!${NC}"
    else
        echo -e "${RED}Database setup failed! Please check your DBUSER and DBPASS.${NC}"
        exit 1
    fi
}

# Main script logic
case "$1" in
    build)
        build
        ;;
    install)
        install
        ;;
    run)
        shift
        run "$@"
        ;;
    setup-db)
        setup_db
        ;;
    uninstall)
        uninstall
        ;;
    help | *)
        usage
        ;;
esac
