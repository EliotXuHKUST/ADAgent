from fastapi import APIRouter

router = APIRouter()

@router.get("/")
def get_ads(user_id: str):
    """
    Get targeted ads for a user
    """
    return {"user_id": user_id, "ads": []}
