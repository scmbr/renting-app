from fastapi_users import schemas


class UserRead(schemas.BaseUser[int]):
    userName: str
    profilePic: str
    phoneNumber: str
    rating: float


class UserCreate(schemas.BaseUserCreate):
    userName: str
    profilePic: str
    phoneNumber: str


class UserUpdate(schemas.BaseUserUpdate):
    userName: str
    profilePic: str
    phoneNumber: str
