from fastapi import APIRouter

router = APIRouter()

@router.post("/events")
def collect_event():
    """
    Collect user behavior events (play, click, search, etc.)
    """
    return {"message": "Event collected"}
