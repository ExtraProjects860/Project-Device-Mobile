from fastapi import APIRouter

router: APIRouter = APIRouter(prefix="/healthcheck", tags=["healthcheck"])


@router.get("/")
def read_root():
    return {"Hello": "World"}
