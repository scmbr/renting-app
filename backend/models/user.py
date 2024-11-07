from sqlalchemy import Column, String, TIMESTAMP, text
from sqlalchemy.orm import Mapped, mapped_column

from db.db import Base


class Users(Base):
    __tablename__ = "users"
    user_name = Column(String(150), nullable=False)
    email = Column(String(255), unique=True, index=True, nullable=False)
    password = Column(String(100), nullable=False)
    status = Column(String, server_default=text("'active'"), nullable=False)
    verified_at = Column(TIMESTAMP(timezone=True))
    created_at = Column(
        TIMESTAMP(timezone=True), server_default=text("now()"), nullable=False
    )
    updated_at = Column(
        TIMESTAMP(timezone=True),
        server_default=text("now()"),
        onupdate=text("now()"),
        nullable=False,
    )
