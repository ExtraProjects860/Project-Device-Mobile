from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

from routers import email, health_check

app = FastAPI()
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

base_path = "/api/v1"

app.include_router(router=email.router, prefix=base_path)
app.include_router(router=health_check.router, prefix=base_path)