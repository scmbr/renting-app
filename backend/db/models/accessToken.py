from fastapi import Depends
from fastapi_users_db_sqlalchemy.access_token import (
    SQLAlchemyBaseAccessTokenTable,
    SQLAlchemyAccessTokenDatabase,
)
from sqlalchemy import Column, Integer, ForeignKey
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.orm import relationship

from db import Base


class AccessToken(SQLAlchemyBaseAccessTokenTable[int], Base):
    __tablename__ = "accesstokens"
    userId = Column(Integer, ForeignKey("users.id"), nullable=False)
    user = relationship("User", back_populates="accessTokens")

    @classmethod
    def get_db(cls, session: "AsyncSession"):
        return SQLAlchemyAccessTokenDatabase(session, cls)
