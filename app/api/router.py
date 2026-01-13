from fastapi import APIRouter
from app.api.endpoints import collection, profile, recommendation, ads

api_router = APIRouter()

api_router.include_router(collection.router, prefix="/collection", tags=["collection"])
api_router.include_router(profile.router, prefix="/profile", tags=["profile"])
api_router.include_router(recommendation.router, prefix="/recommendation", tags=["recommendation"])
api_router.include_router(ads.router, prefix="/ads", tags=["ads"])
