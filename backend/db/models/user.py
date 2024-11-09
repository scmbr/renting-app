from sqlalchemy import Column, String, TIMESTAMP, text, Integer, Float
from sqlalchemy.orm import Mapped, mapped_column
from fastapi_users.db import SQLAlchemyBaseUserTable, SQLAlchemyUserDatabase
from db.base import Base


class User(SQLAlchemyBaseUserTable[int], Base):
    __tablename__ = "users"
    userId = Column(Integer, primary_key=True)
    userName = Column(String(50), nullable=False)
    profilePic = Column(String)
    phoneNumber = Column(String(10))
    rating = Column(Float)
