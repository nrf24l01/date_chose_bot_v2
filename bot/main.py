import asyncio
import logging
import sys
from config import BOT_TOKEN, WEBAPP_URL, BACKEND_URL

from aiogram import Bot, Dispatcher
from aiogram.client.default import DefaultBotProperties
from aiogram.enums import ParseMode
from aiogram.filters import CommandStart
from aiogram.types import Message, CallbackQuery
from aiogram.types import InlineKeyboardButton, InlineKeyboardMarkup, WebAppInfo

from backend_api import BackendAPI

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
    dates = await backend_api.get_user_choices(callback_query.from_user.id)
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


async def main() -> None:
    global backend_api
    async with BackendAPI(base_url=BACKEND_URL) as api:
        backend_api = api  # теперь это глобальное присваивание

        bot = Bot(token=BOT_TOKEN, default=DefaultBotProperties(parse_mode=ParseMode.HTML))
        await dp.start_polling(bot)


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO, stream=sys.stdout)
    asyncio.run(main())
