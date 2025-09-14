from fastapi import FastAPI
from uvicorn import Config, Server
from config.Env import Env

from routers import email, health_check

app = FastAPI()
base_path = "/api/v1"

app.include_router(router=email.router, prefix=base_path)
app.include_router(router=health_check.router, prefix=base_path)

if __name__ == "__main__":
    Server(
        Config("main:app", port=Env.API_PORT, log_level="debug", reload=True)
    ).run()