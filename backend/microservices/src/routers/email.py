from fastapi import APIRouter

router: APIRouter = APIRouter(prefix="/email", tags=["email"])


@router.get("/")
def read_root():
    return {"Hello": "World"}
