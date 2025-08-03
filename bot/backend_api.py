import aiohttp
from typing import List, Optional


class BackendAPI:
    def __init__(self, base_url: str):
        self.base_url = base_url
        self._session: Optional[aiohttp.ClientSession] = None

    async def __aenter__(self):
        self._session = aiohttp.ClientSession()
        return self

    async def __aexit__(self, exc_type, exc_val, exc_tb):
        if self._session:
            await self._session.close()

    async def get_users_votes(self) -> List[dict]:
        if not self._session:
            raise RuntimeError("Session not initialized")
        async with self._session.get(f"{self.base_url}/date/all") as response:
            response.raise_for_status()
            return await response.json()

    async def get_user_choice(self, user_id: int) -> Optional[List[str]]:
        users_votes = await self.get_users_votes()
        for user_vote in users_votes:
            if user_vote.get("user_id") == user_id:
                return user_vote.get("dates")
        return None
