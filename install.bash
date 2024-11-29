#!/bin/bash




SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_DIR="$./TO_DO_CLI" 


GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' 


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


run() {
    cd "$PROJECT_DIR" || exit 1
    go run ./ "$@"
}


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

setup_db() {
    echo -e "${GREEN}Setting up the database...${NC}"

    # Ensure the DBUSER and DBPASS are set
    if [ -z "$DBUSER" ] || [ -z "$DBPASS" ]; then
        echo -e "${RED}DBUSER and DBPASS environment variables are not set!${NC}"
        echo -e "Please set them using:\nexport DBUSER=<your_db_user>\nexport DBPASS=<your_db_password>"
        exit 1
    fi


    SQL=$(cat <<EOF
USE TODOS ; 
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


    echo "$SQL" | mysql -u"$DBUSER" -p"$DBPASS"
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}Database setup completed successfully!${NC}"
    else
        echo -e "${RED}Database setup failed! Please check your DBUSER and DBPASS.${NC}"
        exit 1
    fi
}


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
