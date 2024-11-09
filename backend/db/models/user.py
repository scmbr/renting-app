from sqlalchemy import Column, String, TIMESTAMP, text, Integer, Float
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.orm import Mapped, mapped_column, relationship
from fastapi_users.db import SQLAlchemyBaseUserTable, SQLAlchemyUserDatabase
from db.base import Base


class User(SQLAlchemyBaseUserTable[int], Base):
    __tablename__ = "users"
    userId = Column(Integer, primary_key=True)
    userName = Column(String(50), nullable=False)
    profilePic = Column(String)
    phoneNumber = Column(String(10))
    rating = Column(Float)
    accessTokens = relationship("AccessToken", back_populates="user")

    @classmethod
    def get_db(cls, session: "AsyncSession"):
        return SQLAlchemyUserDatabase(session, cls)
