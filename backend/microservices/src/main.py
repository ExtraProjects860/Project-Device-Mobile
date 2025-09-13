from fastapi import FastAPI

from routers import email, health_check

app = FastAPI()
base_path = "/api/v1"

app.include_router(router=email.router, prefix=base_path)
app.include_router(router=health_check.router, prefix=base_path)
