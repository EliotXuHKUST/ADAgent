from fastapi import APIRouter

router = APIRouter()

@router.get("/{user_id}")
def get_user_profile(user_id: str):
    """
    Get user profile by user_id
    """
    return {"user_id": user_id, "tags": []}
