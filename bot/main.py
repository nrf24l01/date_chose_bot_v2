import asyncio
import logging
import sys
from config import BOT_TOKEN, WEBAPP_URL, BACKEND_URL, ADMIN_ID

from aiogram import Bot, Dispatcher
from aiogram.client.default import DefaultBotProperties
from aiogram.enums import ParseMode
from aiogram.filters import CommandStart
from aiogram.types import Message, CallbackQuery
from aiogram.types import InlineKeyboardButton, InlineKeyboardMarkup, WebAppInfo

from backend_api import BackendAPI
from collections import Counter, defaultdict

dp = Dispatcher()

backend_api: BackendAPI | None = None  # глобально объявленная переменная


@dp.message(CommandStart())
async def command_start_handler(message: Message) -> None:
    keyboard = InlineKeyboardMarkup(inline_keyboard=[
        [
            InlineKeyboardButton(
                text="Выбор дат",
                web_app=WebAppInfo(url=WEBAPP_URL)
            ),
            InlineKeyboardButton(
                text="Я выбрал",
                callback_data="confirm"
            )
        ]
    ])
    await message.answer("Выберите даты в интерфейсе по кнопочке снизу.", reply_markup=keyboard)


@dp.callback_query(lambda c: c.data == "confirm")
async def confirm_callback_handler(callback_query: CallbackQuery) -> None:
    global backend_api
    assert backend_api is not None, "backend_api не инициализирован"
    dates = await backend_api.get_user_choice(callback_query.from_user.id)
    if not dates:
        new_text = "Вы не выбрали даты, хватит пытаться обмануть меня."
        # Compare both text and reply_markup to avoid TelegramBadRequest
        current_text = callback_query.message.text
        current_markup = callback_query.message.reply_markup
        if current_text != new_text or current_markup != callback_query.message.reply_markup:
            await callback_query.message.edit_text(
                new_text,
                reply_markup=callback_query.message.reply_markup
            )
        return
    if len(dates) > 0 and dates[0] != "0001-01-01":
        new_text = f"Вы выбрали даты. Спасибо!\nВаши выбранные даты: {', '.join(dates)}.\nВы можете их изменить, открой веб приложение заново по кнопке."
    else:
        new_text = f"Вы выбрали даты(ну почти). Спасибо!\nНо лучше бы Вы смогли прийти.\nВы можете их изменить, открой веб приложение заново по кнопке."
    current_text = callback_query.message.text
    current_markup = callback_query.message.reply_markup
    if current_text != new_text or current_markup != callback_query.message.reply_markup:
        await callback_query.message.edit_text(
            new_text,
            reply_markup=callback_query.message.reply_markup
        )

@dp.message(lambda message: message.text == "/votes")
async def votes_command_handler(message: Message) -> None:
    if message.from_user.id != ADMIN_ID:
        await message.answer("ПАСИБА за помощь в тестировании, но это уже проверено.")
        return
    global backend_api
    assert backend_api is not None, "backend_api не инициализирован"
    votes = await backend_api.get_users_votes()
    date_votes = defaultdict(list)
    all_users = []
    voted_users = set()
    for vote in votes:
        user_link = (
            f'<a href="tg://user?id={vote["user_id"]}">{vote["user_name"] or vote["user_id"]}</a>'
        )
        all_users.append(user_link)
        has_vote = False
        for date in vote["dates"]:
            if date != "0001-01-01":
                date_votes[date].append(user_link)
                has_vote = True
        if has_vote:
            voted_users.add(user_link)

    not_voted_users = [user for user in all_users if user not in voted_users]

    if not date_votes:
        await message.answer("Пока никто не выбрал даты.")
    else:
        counts = Counter({date: len(users) for date, users in date_votes.items()})
        lines = []
        for date, count in counts.most_common():
            users = ", ".join(date_votes[date])
            lines.append(f"{date} - {count} голосов - {users}")
        if not_voted_users:
            lines.append("\nНе проголосовали:\n" + ", ".join(not_voted_users))
        lines.append("\nВсего пользователей: " + ", ".join(all_users) + ". Количество: " + len(all_users))
        lines.append("Нахрена те эта инфа, ты же знаешь за что дёргать чтоб получить нужную жсонку, сам же совершил осмысленый выбор забить нахрен на безопасность данных, тк ставишь на то что никто не полезет на гитхаб искать исходники это хрени, копаться в них чтоб найти в гойском коде эту ручку или свагеровские ямлы читать, а потом ведь ещё бы пришлось лезть в ci/cd, смотреть как там этот экшн собирает этот долбаный фронт, как витьку засовывается VITE_BACKEND_URL и тд. И вообще, нахрена я это пишу, как будто в 3 часа ночи надо под кровать смотреть(если бы там не было ящиков), а не тыкать на кнопочки на компухтерах своих.")
        await message.answer("\n".join(lines), parse_mode="HTML")

async def main() -> None:
    global backend_api
    async with BackendAPI(base_url=BACKEND_URL) as api:
        backend_api = api  # теперь это глобальное присваивание

        bot = Bot(token=BOT_TOKEN, default=DefaultBotProperties(parse_mode=ParseMode.HTML))
        await dp.start_polling(bot)


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO, stream=sys.stdout)
    asyncio.run(main())
