from dotenv import load_dotenv
from os import getenv

load_dotenv()

# TG config
BOT_TOKEN = getenv("BOT_TOKEN")
WEBAPP_URL = getenv("WEBAPP_URL")
BACKEND_URL = getenv("BACKEND_URL")
ADMIN_ID = int(getenv("ADMIN_ID", 0))