package seeders

import (
	"Printer3DCourses/internal/models"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

func SeedTestsAndQuestionsForCourse1(db *gorm.DB) error {
	// Создаем тест 1: «Основы 3D-печати»
	test1 := models.Test{
		ID:        1,
		TestTitle: "Основы 3D-печати",
		CourseID:  1, // Предположим, что курс с ID = 1 уже существует
	}

	// Создаем тест 2: «Безопасность при 3D-печати»
	test2 := models.Test{
		ID:        2,
		TestTitle: "Безопасность при 3D-печати",
		CourseID:  1, // Предположим, что курс с ID = 1 уже существует
	}

	// Добавляем тесты в базу данных
	if err := db.FirstOrCreate(&test1).Error; err != nil {
		log.Fatal("Ошибка при создании теста 1: ", err)
		return err
	}
	if err := db.FirstOrCreate(&test2).Error; err != nil {
		log.Fatal("Ошибка при создании теста 2: ", err)
		return err
	}

	// Вопросы и ответы для теста 1: «Основы 3D-печати»
	questionsTest1 := []models.Question{
		{
			ID:           1,
			QuestionText: "Какой материал чаще всего используется в FDM-печати?",
			TestID:       test1.ID,
			Answers: []models.Answer{
				{ID: 1, AnswerText: "PLA", IsCorrect: true, QuestionID: 1},
				{ID: 2, AnswerText: "ABS", IsCorrect: false, QuestionID: 1},
				{ID: 3, AnswerText: "Нейлон", IsCorrect: false, QuestionID: 1},
				{ID: 4, AnswerText: "Резина", IsCorrect: false, QuestionID: 1},
			},
		},
		{
			ID:           2,
			QuestionText: "Что означает аббревиатура FDM?",
			TestID:       test1.ID,
			Answers: []models.Answer{
				{ID: 5, AnswerText: "Fused Deposition Modeling", IsCorrect: true, QuestionID: 2},
				{ID: 6, AnswerText: "Flexible Direct Manufacturing", IsCorrect: false, QuestionID: 2},
				{ID: 7, AnswerText: "Fast Digital Modeling", IsCorrect: false, QuestionID: 2},
			},
		},
		{
			ID:           3,
			QuestionText: "Какой принтер подходит для печати мелкими деталями?",
			TestID:       test1.ID,
			Answers: []models.Answer{
				{ID: 8, AnswerText: "FDM", IsCorrect: false, QuestionID: 3},
				{ID: 9, AnswerText: "SLA", IsCorrect: true, QuestionID: 3},
				{ID: 10, AnswerText: "SLS", IsCorrect: false, QuestionID: 3},
			},
		},
		{
			ID:           4,
			QuestionText: "Какая температура платформы для PLA?",
			TestID:       test1.ID,
			Answers: []models.Answer{
				{ID: 11, AnswerText: "50–60°C", IsCorrect: true, QuestionID: 4},
				{ID: 12, AnswerText: "100–120°C", IsCorrect: false, QuestionID: 4},
				{ID: 13, AnswerText: "Не нужен подогрев", IsCorrect: false, QuestionID: 4},
			},
		},
		{
			ID:           5,
			QuestionText: "Что делает слайсер?",
			TestID:       test1.ID,
			Answers: []models.Answer{
				{ID: 14, AnswerText: "Конвертирует 3D-модель в инструкции для принтера", IsCorrect: true, QuestionID: 5},
				{ID: 15, AnswerText: "Рисует 3D-модели", IsCorrect: false, QuestionID: 5},
				{ID: 16, AnswerText: "Чинит принтеры", IsCorrect: false, QuestionID: 5},
			},
		},
		{
			ID:           6,
			QuestionText: "Какой материал токсичен при печати?",
			TestID:       test1.ID,
			Answers: []models.Answer{
				{ID: 17, AnswerText: "PLA", IsCorrect: false, QuestionID: 6},
				{ID: 18, AnswerText: "ABS", IsCorrect: true, QuestionID: 6},
				{ID: 19, AnswerText: "PETG", IsCorrect: false, QuestionID: 6},
			},
		},
		{
			ID:           7,
			QuestionText: "Что такое «перекос слоёв»?",
			TestID:       test1.ID,
			Answers: []models.Answer{
				{ID: 20, AnswerText: "Ошибка печати из-за рассинхронизации осей", IsCorrect: true, QuestionID: 7},
				{ID: 21, AnswerText: "Вид 3D-моделирования", IsCorrect: false, QuestionID: 7},
				{ID: 22, AnswerText: "Тип принтера", IsCorrect: false, QuestionID: 7},
			},
		},
		{
			ID:           8,
			QuestionText: "Для чего нужна поддержка (supports)?",
			TestID:       test1.ID,
			Answers: []models.Answer{
				{ID: 23, AnswerText: "Чтобы печатать нависающие элементы", IsCorrect: true, QuestionID: 8},
				{ID: 24, AnswerText: "Для ускорения печати", IsCorrect: false, QuestionID: 8},
				{ID: 25, AnswerText: "Для украшения модели", IsCorrect: false, QuestionID: 8},
			},
		},
	}

	// Вставляем вопросы и ответы для теста 1 в базу данных
	if err := db.FirstOrCreate(&questionsTest1).Error; err != nil {
		log.Fatal("Ошибка при добавлении вопросов теста 1: ", err)
		return err
	}

	// Вопросы и ответы для теста 2: «Безопасность при 3D-печати»
	questionsTest2 := []models.Question{
		{
			ID:           9,
			QuestionText: "Можно ли оставлять 3D-принтер без присмотра?",
			TestID:       test2.ID,
			Answers: []models.Answer{
				{ID: 26, AnswerText: "Да, если печатает PLA", IsCorrect: false, QuestionID: 9},
				{ID: 27, AnswerText: "Нет, всегда нужно контролировать процесс", IsCorrect: true, QuestionID: 9},
			},
		},
		{
			ID:           10,
			QuestionText: "Какая температура экструдера для PLA?",
			TestID:       test2.ID,
			Answers: []models.Answer{
				{ID: 28, AnswerText: "190–220°C", IsCorrect: true, QuestionID: 10},
				{ID: 29, AnswerText: "250–300°C", IsCorrect: false, QuestionID: 10},
				{ID: 30, AnswerText: "150–170°C", IsCorrect: false, QuestionID: 10},
			},
		},
		{
			ID:           11,
			QuestionText: "Что делать при засорении сопла?",
			TestID:       test2.ID,
			Answers: []models.Answer{
				{ID: 31, AnswerText: "Очистить иглой или холодной протяжкой", IsCorrect: true, QuestionID: 11},
				{ID: 32, AnswerText: "Стучать по нему молотком", IsCorrect: false, QuestionID: 11},
				{ID: 33, AnswerText: "Залить водой", IsCorrect: false, QuestionID: 11},
			},
		},
		{
			ID:           12,
			QuestionText: "Нужна ли вентиляция при печати PLA?",
			TestID:       test2.ID,
			Answers: []models.Answer{
				{ID: 34, AnswerText: "Да, обязательна", IsCorrect: false, QuestionID: 12},
				{ID: 35, AnswerText: "Желательна, но не критична", IsCorrect: true, QuestionID: 12},
				{ID: 36, AnswerText: "Не нужна", IsCorrect: false, QuestionID: 12},
			},
		},
		{
			ID:           13,
			QuestionText: "Можно ли трогать печатную платформу во время работы?",
			TestID:       test2.ID,
			Answers: []models.Answer{
				{ID: 37, AnswerText: "Да, если в перчатках", IsCorrect: false, QuestionID: 13},
				{ID: 38, AnswerText: "Нет, это опасно!", IsCorrect: true, QuestionID: 13},
				{ID: 39, AnswerText: "Только если принтер выключен", IsCorrect: false, QuestionID: 13},
			},
		},
		{
			ID:           14,
			QuestionText: "Что делать, если модель не прилипает к столу?",
			TestID:       test2.ID,
			Answers: []models.Answer{
				{ID: 40, AnswerText: "Проверить калибровку и использовать клей-карандаш", IsCorrect: true, QuestionID: 14},
				{ID: 41, AnswerText: "Увеличить скорость печати", IsCorrect: false, QuestionID: 14},
				{ID: 42, AnswerText: "Ничего, само прилипнет", IsCorrect: false, QuestionID: 14},
			},
		},
		{
			ID:           15,
			QuestionText: "Какой инструмент нужен для снятия модели?",
			TestID:       test2.ID,
			Answers: []models.Answer{
				{ID: 43, AnswerText: "Шпатель", IsCorrect: true, QuestionID: 15},
				{ID: 44, AnswerText: "Ножницы", IsCorrect: false, QuestionID: 15},
				{ID: 45, AnswerText: "Отвёртка", IsCorrect: false, QuestionID: 15},
			},
		},
		{
			ID:           16,
			QuestionText: "Куда девать опоры после печати?",
			TestID:       test2.ID,
			Answers: []models.Answer{
				{ID: 46, AnswerText: "Аккуратно отломать или срезать", IsCorrect: true, QuestionID: 16},
				{ID: 47, AnswerText: "Растворить в воде", IsCorrect: false, QuestionID: 16},
				{ID: 48, AnswerText: "Оставить на модели", IsCorrect: false, QuestionID: 16},
			},
		},
	}

	// Вставляем вопросы и ответы для теста 2 в базу данных
	if err := db.FirstOrCreate(&questionsTest2).Error; err != nil {
		log.Fatal("Ошибка при добавлении вопросов теста 2: ", err)
		return err
	}

	return nil
}

func SeedTestsAndQuestionsForCourse2(db *gorm.DB) error {
	// Тест 1: Основы работы в TinkerCAD
	test1 := models.Test{
		ID:        3,
		TestTitle: "Основы работы в TinkerCAD",
		CourseID:  2,
	}

	if err := db.FirstOrCreate(&test1, models.Test{ID: test1.ID}).Error; err != nil {
		return err
	}

	questions1 := []models.Question{
		{ID: 17, QuestionText: "Какой инструмент используется для создания базовых фигур в TinkerCAD?", TestID: test1.ID},
		{ID: 18, QuestionText: "Как изменить размер объекта в TinkerCAD?", TestID: test1.ID},
		{ID: 19, QuestionText: "Как создать отверстие в объекте?", TestID: test1.ID},
		{ID: 20, QuestionText: "Как повернуть объект в 3D-пространстве?", TestID: test1.ID},
		{ID: 21, QuestionText: "Как сохранить проект в TinkerCAD?", TestID: test1.ID},
		{ID: 22, QuestionText: "Как продублировать объект?", TestID: test1.ID},
		{ID: 23, QuestionText: "Как выровнять несколько объектов относительно друг друга?", TestID: test1.ID},
		{ID: 24, QuestionText: "Какой формат файла можно экспортировать для 3D-печати?", TestID: test1.ID},
	}

	for _, q := range questions1 {
		db.FirstOrCreate(&q, models.Question{ID: q.ID, TestID: test1.ID})
	}

	answers1 := []models.Answer{
		{ID: 49, AnswerText: "Инструмент \"Основные фигуры\"", IsCorrect: true, QuestionID: 17},
		{ID: 50, AnswerText: "Инструмент \"Карандаш\"", IsCorrect: false, QuestionID: 17},
		{ID: 51, AnswerText: "Инструмент \"Кривые Безье\"", IsCorrect: false, QuestionID: 17},
		{ID: 52, AnswerText: "Инструмент \"Текст\"", IsCorrect: false, QuestionID: 17},

		{ID: 53, AnswerText: "Перетаскивание белых маркеров по углам", IsCorrect: true, QuestionID: 18},
		{ID: 54, AnswerText: "Двойной щелчок по объекту", IsCorrect: false, QuestionID: 18},
		{ID: 55, AnswerText: "Через меню \"Правка → Масштаб\"", IsCorrect: false, QuestionID: 18},
		{ID: 56, AnswerText: "Нажатием клавиш Ctrl + S", IsCorrect: false, QuestionID: 18},

		{ID: 57, AnswerText: "Преобразовать объект в \"Отверстие\" и сгруппировать с основным", IsCorrect: true, QuestionID: 19},
		{ID: 58, AnswerText: "Инструментом \"Нож\"", IsCorrect: false, QuestionID: 19},
		{ID: 59, AnswerText: "Удалить часть объекта ластиком", IsCorrect: false, QuestionID: 19},
		{ID: 60, AnswerText: "Использовать функцию \"Вырезать\"", IsCorrect: false, QuestionID: 19},

		{ID: 61, AnswerText: "Используя круговые стрелки вокруг объекта", IsCorrect: true, QuestionID: 20},
		{ID: 62, AnswerText: "Через меню \"Вид → Поворот\"", IsCorrect: false, QuestionID: 20},
		{ID: 63, AnswerText: "Комбинацией клавиш Ctrl + R", IsCorrect: false, QuestionID: 20},
		{ID: 64, AnswerText: "Инструментом \"Трансформация\"", IsCorrect: false, QuestionID: 20},

		{ID: 65, AnswerText: "Автоматически сохраняется в облако", IsCorrect: true, QuestionID: 21},
		{ID: 66, AnswerText: "Файл → Сохранить как STL", IsCorrect: false, QuestionID: 21},
		{ID: 67, AnswerText: "Нажать кнопку \"Экспорт\"", IsCorrect: false, QuestionID: 21},
		{ID: 68, AnswerText: "Через настройки аккаунта", IsCorrect: false, QuestionID: 21},

		{ID: 69, AnswerText: "Ctrl + D или кнопка \"Дублировать\"", IsCorrect: true, QuestionID: 22},
		{ID: 70, AnswerText: "Правка → Копировать", IsCorrect: false, QuestionID: 22},
		{ID: 71, AnswerText: "Перетащить с зажатой правой кнопкой мыши", IsCorrect: false, QuestionID: 22},
		{ID: 72, AnswerText: "Через контекстное меню", IsCorrect: false, QuestionID: 22},

		{ID: 73, AnswerText: "Инструмент \"Выровнять\" на панели инструментов", IsCorrect: true, QuestionID: 23},
		{ID: 74, AnswerText: "Вручную перемещать по сетке", IsCorrect: false, QuestionID: 23},
		{ID: 75, AnswerText: "Через меню \"Группировка\"", IsCorrect: false, QuestionID: 23},
		{ID: 76, AnswerText: "Автоматически при близком расположении", IsCorrect: false, QuestionID: 23},

		{ID: 77, AnswerText: ".STL", IsCorrect: true, QuestionID: 24},
		{ID: 78, AnswerText: ".OBJ", IsCorrect: false, QuestionID: 24},
		{ID: 79, AnswerText: ".GCODE", IsCorrect: false, QuestionID: 24},
		{ID: 80, AnswerText: ".STEP", IsCorrect: false, QuestionID: 24},
	}

	for _, ans := range answers1 {
		db.FirstOrCreate(&ans, models.Answer{ID: ans.ID, QuestionID: ans.QuestionID})
	}

	// Тест 2: Практическое моделирование в TinkerCAD
	test2 := models.Test{
		ID:        4,
		TestTitle: "Практическое моделирование в TinkerCAD",
		CourseID:  2,
	}

	if err := db.FirstOrCreate(&test2, models.Test{ID: test2.ID}).Error; err != nil {
		return err
	}

	questions2 := []models.Question{
		{ID: 25, QuestionText: "Как создать надпись на 3D-объекте?", TestID: test2.ID},
		{ID: 26, QuestionText: "Как изменить цвет объекта?", TestID: test2.ID},
		{ID: 27, QuestionText: "Как создать сложную фигуру из нескольких объектов?", TestID: test2.ID},
		{ID: 28, QuestionText: "Как переместить объект по оси Z (вверх/вниз)?", TestID: test2.ID},
		{ID: 29, QuestionText: "Для чего нужна функция \"Скрыть объект\"?", TestID: test2.ID},
		{ID: 30, QuestionText: "Как добавить готовый компонент (например, Arduino)?", TestID: test2.ID},
		{ID: 31, QuestionText: "Как проверить модель на ошибки перед печатью?", TestID: test2.ID},
		{ID: 32, QuestionText: "Как создать полый объект (например, коробку)?", TestID: test2.ID},
	}

	for _, q := range questions2 {
		db.FirstOrCreate(&q, models.Question{ID: q.ID, TestID: test2.ID})
	}

	// Продолжить с ID 81+
	answers2 := []models.Answer{
		{ID: 81, AnswerText: "Добавить объект \"Текст\" и сгруппировать с основным", IsCorrect: true, QuestionID: 25},
		{ID: 82, AnswerText: "Инструментом \"Гравировка\"", IsCorrect: false, QuestionID: 25},
		{ID: 83, AnswerText: "Импортировать изображение", IsCorrect: false, QuestionID: 25},
		{ID: 84, AnswerText: "Через меню \"Эффекты\"", IsCorrect: false, QuestionID: 25},

		{ID: 85, AnswerText: "Выбрать объект и изменить цвет в палитре", IsCorrect: true, QuestionID: 26},
		{ID: 86, AnswerText: "Инструментом \"Заливка\"", IsCorrect: false, QuestionID: 26},
		{ID: 87, AnswerText: "Через настройки материала", IsCorrect: false, QuestionID: 26},
		{ID: 88, AnswerText: "Только при экспорте", IsCorrect: false, QuestionID: 26},

		{ID: 89, AnswerText: "Сгруппировать объекты (Ctrl + G)", IsCorrect: true, QuestionID: 27},
		{ID: 90, AnswerText: "Инструментом \"Булевы операции\"", IsCorrect: false, QuestionID: 27},
		{ID: 91, AnswerText: "Использовать \"Объединить\" в меню", IsCorrect: false, QuestionID: 27},
		{ID: 92, AnswerText: "Вручную совместить грани", IsCorrect: false, QuestionID: 27},

		{ID: 93, AnswerText: "Черной стрелкой над объектом", IsCorrect: true, QuestionID: 28},
		{ID: 94, AnswerText: "Клавишами Q и E", IsCorrect: false, QuestionID: 28},
		{ID: 95, AnswerText: "Через параметры высоты", IsCorrect: false, QuestionID: 28},
		{ID: 96, AnswerText: "Только в режиме \"Редактирование сетки\"", IsCorrect: false, QuestionID: 28},

		{ID: 97, AnswerText: "Временно убрать объект из поля зрения", IsCorrect: true, QuestionID: 29},
		{ID: 98, AnswerText: "Удалить объект без возможности восстановления", IsCorrect: false, QuestionID: 29},
		{ID: 99, AnswerText: "Сделать объект прозрачным", IsCorrect: false, QuestionID: 29},
		{ID: 100, AnswerText: "Экспортировать только видимые части", IsCorrect: false, QuestionID: 29},

		{ID: 101, AnswerText: "Выбрать из библиотеки \"Компоненты\"", IsCorrect: true, QuestionID: 30},
		{ID: 102, AnswerText: "Импортировать из файла", IsCorrect: false, QuestionID: 30},
		{ID: 103, AnswerText: "Скачать с официального сайта", IsCorrect: false, QuestionID: 30},
		{ID: 104, AnswerText: "Использовать функцию \"Дублировать\"", IsCorrect: false, QuestionID: 30},

		{ID: 105, AnswerText: "Использовать инструмент \"Исправить ошибки\"", IsCorrect: true, QuestionID: 31},
		{ID: 106, AnswerText: "Процесс идет автоматически", IsCorrect: false, QuestionID: 31},
		{ID: 107, AnswerText: "Модель нужно распечатать для проверки", IsCorrect: false, QuestionID: 31},
		{ID: 108, AnswerText: "Проверять вручную", IsCorrect: false, QuestionID: 31},

		{ID: 109, AnswerText: "Инструментом \"Полый\"", IsCorrect: true, QuestionID: 32},
		{ID: 110, AnswerText: "Использовать \"Прочистить\"", IsCorrect: false, QuestionID: 32},
		{ID: 111, AnswerText: "Инструментом \"Объемный вырез\"", IsCorrect: false, QuestionID: 32},
		{ID: 112, AnswerText: "Включить опцию \"Пустота\" в настройках объекта", IsCorrect: false, QuestionID: 32},
	}

	for _, ans := range answers2 {
		db.FirstOrCreate(&ans, models.Answer{ID: ans.ID, QuestionID: ans.QuestionID})
	}

	return nil
}

func SeedTestsAndQuestionsForCourse3(db *gorm.DB) error {
	// Тест 1: Основы калибровки и настройки
	test1 := models.Test{
		ID:        5,
		TestTitle: "Основы калибровки и настройки",
		CourseID:  3,
	}

	if err := db.FirstOrCreate(&test1, models.Test{ID: test1.ID}).Error; err != nil {
		return err
	}

	questions1 := []models.Question{
		{ID: 33, QuestionText: "Какой первый шаг в калибровке 3D-принтера?", TestID: test1.ID},
		{ID: 34, QuestionText: "Как проверить, что сопло находится на правильном расстоянии от стола?", TestID: test1.ID},
		{ID: 35, QuestionText: "Что регулируется при калибровке \"шагов на мм\" для оси E (экструдера)?", TestID: test1.ID},
		{ID: 36, QuestionText: "Как исправить \"эффект слоёвного смещения\" (layer shifting)?", TestID: test1.ID},
		{ID: 37, QuestionText: "Для чего нужна авто-калибровка стола (BLTouch, CR-Touch)?", TestID: test1.ID},
		{ID: 38, QuestionText: "Какая температура стола рекомендуется для PLA?", TestID: test1.ID},
		{ID: 39, QuestionText: "Какой параметр влияет на прочность модели?", TestID: test1.ID},
		{ID: 40, QuestionText: "Что делать, если первая печать не прилипает к столу?", TestID: test1.ID},
	}

	for _, q := range questions1 {
		db.FirstOrCreate(&q, models.Question{ID: q.ID, TestID: test1.ID})
	}

	answers1 := []models.Answer{
		{ID: 113, AnswerText: "Выравнивание стола (калибровка оси Z)", IsCorrect: true, QuestionID: 33},
		{ID: 114, AnswerText: "Настройка температуры экструдера", IsCorrect: false, QuestionID: 33},
		{ID: 115, AnswerText: "Проверка напряжения блока питания", IsCorrect: false, QuestionID: 33},
		{ID: 116, AnswerText: "Калибровка шагов на мм для осей", IsCorrect: false, QuestionID: 33},

		{ID: 117, AnswerText: "Использовать лист бумаги (должен слегка зажиматься)", IsCorrect: true, QuestionID: 34},
		{ID: 118, AnswerText: "Измерить линейкой", IsCorrect: false, QuestionID: 34},
		{ID: 119, AnswerText: "На глаз, без зазора", IsCorrect: false, QuestionID: 34},
		{ID: 120, AnswerText: "Коснуться соплом стола", IsCorrect: false, QuestionID: 34},

		{ID: 121, AnswerText: "Количество подаваемого филамента", IsCorrect: true, QuestionID: 35},
		{ID: 122, AnswerText: "Скорость печати", IsCorrect: false, QuestionID: 35},
		{ID: 123, AnswerText: "Температура нагревательного блока", IsCorrect: false, QuestionID: 35},
		{ID: 124, AnswerText: "Высота слоя", IsCorrect: false, QuestionID: 35},

		{ID: 125, AnswerText: "Проверить натяжение ремней и состояние шкивов", IsCorrect: true, QuestionID: 36},
		{ID: 126, AnswerText: "Увеличить температуру стола", IsCorrect: false, QuestionID: 36},
		{ID: 127, AnswerText: "Заменить филамент", IsCorrect: false, QuestionID: 36},
		{ID: 128, AnswerText: "Уменьшить скорость печати", IsCorrect: false, QuestionID: 36},

		{ID: 129, AnswerText: "Компенсация неровностей поверхности стола", IsCorrect: true, QuestionID: 37},
		{ID: 130, AnswerText: "Измерение температуры сопла", IsCorrect: false, QuestionID: 37},
		{ID: 131, AnswerText: "Контроль влажности филамента", IsCorrect: false, QuestionID: 37},
		{ID: 132, AnswerText: "Настройка скорости вентилятора", IsCorrect: false, QuestionID: 37},

		{ID: 133, AnswerText: "50–60°C", IsCorrect: true, QuestionID: 38},
		{ID: 134, AnswerText: "20–30°C", IsCorrect: false, QuestionID: 38},
		{ID: 135, AnswerText: "100–110°C", IsCorrect: false, QuestionID: 38},
		{ID: 136, AnswerText: "120–130°C", IsCorrect: false, QuestionID: 38},

		{ID: 137, AnswerText: "Заполнение (infill) и количество периметров", IsCorrect: true, QuestionID: 39},
		{ID: 138, AnswerText: "Цвет филамента", IsCorrect: false, QuestionID: 39},
		{ID: 139, AnswerText: "Скорость движения вентилятора", IsCorrect: false, QuestionID: 39},
		{ID: 140, AnswerText: "Яркость подсветки принтера", IsCorrect: false, QuestionID: 39},

		{ID: 141, AnswerText: "Проверить уровень стола, очистить поверхность и использовать адгезив (клей, лак для волос)", IsCorrect: true, QuestionID: 40},
		{ID: 142, AnswerText: "Увеличить скорость печати", IsCorrect: false, QuestionID: 40},
		{ID: 143, AnswerText: "Отключить подогрев стола", IsCorrect: false, QuestionID: 40},
		{ID: 144, AnswerText: "Печатать без подложки", IsCorrect: false, QuestionID: 40},
	}

	for _, ans := range answers1 {
		db.FirstOrCreate(&ans, models.Answer{ID: ans.ID, QuestionID: ans.QuestionID})
	}

	// Тест 2: Решение проблем и тонкие настройки
	test2 := models.Test{
		ID:        6,
		TestTitle: "Решение проблем и тонкие настройки",
		CourseID:  3,
	}

	if err := db.FirstOrCreate(&test2, models.Test{ID: test2.ID}).Error; err != nil {
		return err
	}

	questions2 := []models.Question{
		{ID: 41, QuestionText: "Как устранить \"волочение\" (oozing) филамента при перемещении?", TestID: test2.ID},
		{ID: 42, QuestionText: "Что означает \"термический перегрев\" (heat creep)?", TestID: test2.ID},
		{ID: 43, QuestionText: "Как исправить \"бублик\" (elephant's foot) в нижней части модели?", TestID: test2.ID},
		{ID: 44, QuestionText: "Для чего нужна PID-калибровка нагревателя?", TestID: test2.ID},
		{ID: 45, QuestionText: "Какой параметр слайсера влияет на гладкость поверхности?", TestID: test2.ID},
		{ID: 46, QuestionText: "Что делать, если филамент не подаётся в экструдер?", TestID: test2.ID},
		{ID: 47, QuestionText: "Как избежать \"стружки\" (stringing) между деталями?", TestID: test2.ID},
		{ID: 48, QuestionText: "Как проверить точность размеров напечатанной детали?", TestID: test2.ID},
	}

	for _, q := range questions2 {
		db.FirstOrCreate(&q, models.Question{ID: q.ID, TestID: test2.ID})
	}

	answers2 := []models.Answer{
		{ID: 145, AnswerText: "Настроить ретракцию (втягивание филамента)", IsCorrect: true, QuestionID: 41},
		{ID: 146, AnswerText: "Увеличить температуру сопла", IsCorrect: false, QuestionID: 41},
		{ID: 147, AnswerText: "Отключить охлаждение", IsCorrect: false, QuestionID: 41},
		{ID: 148, AnswerText: "Печатать на высокой скорости", IsCorrect: false, QuestionID: 41},

		{ID: 149, AnswerText: "Нагрев филамента выше рабочей зоны, приводящий к засорению", IsCorrect: true, QuestionID: 42},
		{ID: 150, AnswerText: "Перегрев блока питания", IsCorrect: false, QuestionID: 42},
		{ID: 151, AnswerText: "Перегрев стола", IsCorrect: false, QuestionID: 42},
		{ID: 152, AnswerText: "Плавление шестерни экструдера", IsCorrect: false, QuestionID: 42},

		{ID: 153, AnswerText: "Уменьшить температуру стола или отрегулировать Z-offset", IsCorrect: true, QuestionID: 43},
		{ID: 154, AnswerText: "Увеличить скорость печати", IsCorrect: false, QuestionID: 43},
		{ID: 155, AnswerText: "Отключить вентилятор", IsCorrect: false, QuestionID: 43},
		{ID: 156, AnswerText: "Использовать более толстый слой", IsCorrect: false, QuestionID: 43},

		{ID: 157, AnswerText: "Стабилизация температуры сопла", IsCorrect: true, QuestionID: 44},
		{ID: 158, AnswerText: "Ускорение печати", IsCorrect: false, QuestionID: 44},
		{ID: 159, AnswerText: "Уменьшение шума двигателей", IsCorrect: false, QuestionID: 44},
		{ID: 160, AnswerText: "Калибровка датчика уровня стола", IsCorrect: false, QuestionID: 44},

		{ID: 161, AnswerText: "Высота слоя (layer height)", IsCorrect: true, QuestionID: 45},
		{ID: 162, AnswerText: "Скорость перемещения", IsCorrect: false, QuestionID: 45},
		{ID: 163, AnswerText: "Температура стола", IsCorrect: false, QuestionID: 45},
		{ID: 164, AnswerText: "Яркость подсветки", IsCorrect: false, QuestionID: 45},

		{ID: 165, AnswerText: "Проверить засорение сопла и натяжение шестерни подачи", IsCorrect: true, QuestionID: 46},
		{ID: 166, AnswerText: "Увеличить скорость печати", IsCorrect: false, QuestionID: 46},
		{ID: 167, AnswerText: "Понизить температуру сопла", IsCorrect: false, QuestionID: 46},
		{ID: 168, AnswerText: "Отключить подогрев стола", IsCorrect: false, QuestionID: 46},

		{ID: 169, AnswerText: "Настроить ретракцию и уменьшить температуру сопла", IsCorrect: true, QuestionID: 47},
		{ID: 170, AnswerText: "Увеличить скорость печати", IsCorrect: false, QuestionID: 47},
		{ID: 171, AnswerText: "Печатать без охлаждения", IsCorrect: false, QuestionID: 47},
		{ID: 172, AnswerText: "Использовать более толстый филамент", IsCorrect: false, QuestionID: 47},

		{ID: 173, AnswerText: "Напечатать калибровочный куб и измерить штангенциркулем", IsCorrect: true, QuestionID: 48},
		{ID: 174, AnswerText: "Взвесить модель", IsCorrect: false, QuestionID: 48},
		{ID: 175, AnswerText: "Проверить на глаз", IsCorrect: false, QuestionID: 48},
		{ID: 176, AnswerText: "Сравнить с 3D-моделью на экране", IsCorrect: false, QuestionID: 48},
	}

	for _, ans := range answers2 {
		db.FirstOrCreate(&ans, models.Answer{ID: ans.ID, QuestionID: ans.QuestionID})
	}

	return nil
}

func SeedTestsAndQuestionsForCourse4(db *gorm.DB) error {
	// Создание теста "Основы электроники и компоненты" для курса 4
	test1 := models.Test{ID: 7, TestTitle: "Основы электроники и компоненты", CourseID: 4}
	db.FirstOrCreate(&test1)

	// Вопросы для теста 1
	questions1 := []models.Question{
		{ID: 49, QuestionText: "Что измеряет мультиметр в режиме \"Ω\"?", TestID: test1.ID},
		{ID: 50, QuestionText: "Как обозначается конденсатор на схеме?", TestID: test1.ID},
		{ID: 51, QuestionText: "Какой компонент используется для защиты цепи от перегрузки?", TestID: test1.ID},
		{ID: 52, QuestionText: "Что делает диод?", TestID: test1.ID},
		{ID: 53, QuestionText: "Какой закон описывает зависимость тока от напряжения и сопротивления?", TestID: test1.ID},
		{ID: 54, QuestionText: "Как называется компонент, усиливающий сигнал?", TestID: test1.ID},
		{ID: 55, QuestionText: "Что такое \"земля\" (GND) в электронике?", TestID: test1.ID},
		{ID: 56, QuestionText: "Какой прибор измеряет частоту сигнала?", TestID: test1.ID},
	}
	for _, q := range questions1 {
		db.FirstOrCreate(&q, models.Question{ID: q.ID, TestID: q.TestID})
	}

	// Ответы для теста 1
	answers1 := []models.Answer{
		{ID: 177, AnswerText: "Сопротивление", IsCorrect: true, QuestionID: 49},
		{ID: 178, AnswerText: "Напряжение", IsCorrect: false, QuestionID: 49},
		{ID: 179, AnswerText: "Силу тока", IsCorrect: false, QuestionID: 49},
		{ID: 180, AnswerText: "Ёмкость", IsCorrect: false, QuestionID: 49},
		{ID: 181, AnswerText: "Две параллельные линии", IsCorrect: true, QuestionID: 50},
		{ID: 182, AnswerText: "Прямоугольник", IsCorrect: false, QuestionID: 50},
		{ID: 183, AnswerText: "Зигзаг", IsCorrect: false, QuestionID: 50},
		{ID: 184, AnswerText: "Треугольник", IsCorrect: false, QuestionID: 50},
		{ID: 185, AnswerText: "Предохранитель", IsCorrect: true, QuestionID: 51},
		{ID: 186, AnswerText: "Резистор", IsCorrect: false, QuestionID: 51},
		{ID: 187, AnswerText: "Транзистор", IsCorrect: false, QuestionID: 51},
		{ID: 188, AnswerText: "Диод", IsCorrect: false, QuestionID: 51},
		{ID: 189, AnswerText: "Пропускает ток только в одном направлении", IsCorrect: true, QuestionID: 52},
		{ID: 190, AnswerText: "Усиливает сигнал", IsCorrect: false, QuestionID: 52},
		{ID: 191, AnswerText: "Хранит заряд", IsCorrect: false, QuestionID: 52},
		{ID: 192, AnswerText: "Измеряет напряжение", IsCorrect: false, QuestionID: 52},
		{ID: 193, AnswerText: "Закон Ома (I = U/R)", IsCorrect: true, QuestionID: 53},
		{ID: 194, AnswerText: "Закон Кирхгофа", IsCorrect: false, QuestionID: 53},
		{ID: 195, AnswerText: "Закон Фарадея", IsCorrect: false, QuestionID: 53},
		{ID: 196, AnswerText: "Закон Джоуля-Ленца", IsCorrect: false, QuestionID: 53},
		{ID: 197, AnswerText: "Транзистор", IsCorrect: true, QuestionID: 54},
		{ID: 198, AnswerText: "Резистор", IsCorrect: false, QuestionID: 54},
		{ID: 199, AnswerText: "Конденсатор", IsCorrect: false, QuestionID: 54},
		{ID: 200, AnswerText: "Катушка индуктивности", IsCorrect: false, QuestionID: 54},
		{ID: 201, AnswerText: "Общая точка цепи с нулевым потенциалом", IsCorrect: true, QuestionID: 55},
		{ID: 202, AnswerText: "Источник питания", IsCorrect: false, QuestionID: 55},
		{ID: 203, AnswerText: "Переменное напряжение", IsCorrect: false, QuestionID: 55},
		{ID: 204, AnswerText: "Изолированный провод", IsCorrect: false, QuestionID: 55},
		{ID: 205, AnswerText: "Осциллограф", IsCorrect: true, QuestionID: 56},
		{ID: 206, AnswerText: "Мультиметр", IsCorrect: false, QuestionID: 56},
		{ID: 207, AnswerText: "Логический анализатор", IsCorrect: false, QuestionID: 56},
		{ID: 208, AnswerText: "Генератор сигналов", IsCorrect: false, QuestionID: 56},
	}

	for _, ans := range answers1 {
		db.FirstOrCreate(&ans, models.Answer{ID: ans.ID, QuestionID: ans.QuestionID})
	}

	// Создание второго теста "Цепи и схемотехника"
	test2 := models.Test{ID: 8, TestTitle: "Цепи и схемотехника", CourseID: 4}
	db.FirstOrCreate(&test2)

	// Вопросы для теста 2
	questions2 := []models.Question{
		{ID: 57, QuestionText: "Как соединить два резистора для увеличения общего сопротивления?", TestID: test2.ID},
		{ID: 58, QuestionText: "Что произойдет, если подключить светодиод без резистора?", TestID: test2.ID},
		{ID: 59, QuestionText: "Какой элемент используется для фильтрации высоких частот?", TestID: test2.ID},
		{ID: 60, QuestionText: "Как называется плата для прототипирования без пайки?", TestID: test2.ID},
		{ID: 61, QuestionText: "Что такое ШИМ (PWM)?", TestID: test2.ID},
		{ID: 62, QuestionText: "Как проверить исправность транзистора мультиметром?", TestID: test2.ID},
		{ID: 63, QuestionText: "Для чего нужен операционный усилитель (ОУ)?", TestID: test2.ID},
		{ID: 64, QuestionText: "Что такое \"pull-up резистор\"?", TestID: test2.ID},
	}
	for _, q := range questions2 {
		db.FirstOrCreate(&q, models.Question{ID: q.ID, TestID: q.TestID})
	}

	// Ответы для теста 2
	answers2 := []models.Answer{
		{ID: 209, AnswerText: "Последовательно", IsCorrect: true, QuestionID: 57},
		{ID: 210, AnswerText: "Параллельно", IsCorrect: false, QuestionID: 57},
		{ID: 211, AnswerText: "Через конденсатор", IsCorrect: false, QuestionID: 57},
		{ID: 212, AnswerText: "Через диод", IsCorrect: false, QuestionID: 57},
		{ID: 213, AnswerText: "Он перегорит", IsCorrect: true, QuestionID: 58},
		{ID: 214, AnswerText: "Будет светить ярче", IsCorrect: false, QuestionID: 58},
		{ID: 215, AnswerText: "Не включится", IsCorrect: false, QuestionID: 58},
		{ID: 216, AnswerText: "Начнёт мигать", IsCorrect: false, QuestionID: 58},
		{ID: 217, AnswerText: "Конденсатор", IsCorrect: true, QuestionID: 59},
		{ID: 218, AnswerText: "Резистор", IsCorrect: false, QuestionID: 59},
		{ID: 219, AnswerText: "Катушка индуктивности", IsCorrect: false, QuestionID: 59},
		{ID: 220, AnswerText: "Транзистор", IsCorrect: false, QuestionID: 59},
		{ID: 221, AnswerText: "Breadboard (макетная плата)", IsCorrect: true, QuestionID: 60},
		{ID: 222, AnswerText: "PCB", IsCorrect: false, QuestionID: 60},
		{ID: 223, AnswerText: "Фольгированный текстолит", IsCorrect: false, QuestionID: 60},
		{ID: 224, AnswerText: "Сенсорная панель", IsCorrect: false, QuestionID: 60},
		{ID: 225, AnswerText: "Управление мощностью через изменение ширины импульсов", IsCorrect: true, QuestionID: 61},
		{ID: 226, AnswerText: "Вид переменного тока", IsCorrect: false, QuestionID: 61},
		{ID: 227, AnswerText: "Метод измерения сопротивления", IsCorrect: false, QuestionID: 61},
		{ID: 228, AnswerText: "Тип конденсатора", IsCorrect: false, QuestionID: 61},
		{ID: 229, AnswerText: "В режиме проверки диодов (hFE)", IsCorrect: true, QuestionID: 62},
		{ID: 230, AnswerText: "В режиме измерения напряжения", IsCorrect: false, QuestionID: 62},
		{ID: 231, AnswerText: "В режиме ёмкости", IsCorrect: false, QuestionID: 62},
		{ID: 232, AnswerText: "В режиме частоты", IsCorrect: false, QuestionID: 62},
		{ID: 233, AnswerText: "Для усиления и обработки сигналов", IsCorrect: true, QuestionID: 63},
		{ID: 234, AnswerText: "Для хранения данных", IsCorrect: false, QuestionID: 63},
		{ID: 235, AnswerText: "Для генерации тепла", IsCorrect: false, QuestionID: 63},
		{ID: 236, AnswerText: "Для измерения освещенности", IsCorrect: false, QuestionID: 63},
		{ID: 237, AnswerText: "Резистор, подтягивающий сигнал к +Vcc", IsCorrect: true, QuestionID: 64},
		{ID: 238, AnswerText: "Резистор в цепи питания", IsCorrect: false, QuestionID: 64},
		{ID: 239, AnswerText: "Переменный резистор", IsCorrect: false, QuestionID: 64},
		{ID: 240, AnswerText: "Резистор для измерения тока", IsCorrect: false, QuestionID: 64},
	}

	for _, ans := range answers2 {
		db.FirstOrCreate(&ans, models.Answer{ID: ans.ID, QuestionID: ans.QuestionID})
	}
	test3 := models.Test{ID: 9, TestTitle: "Практическая радиоэлектроника", CourseID: 4}
	db.FirstOrCreate(&test3)

	// Вопросы для теста 3
	questions3 := []models.Question{
		{ID: 65, QuestionText: "Какой инструмент нужен для пайки SMD-компонентов?", TestID: test3.ID},
		{ID: 66, QuestionText: "Что делает преобразователь 'USB-UART'?", TestID: test3.ID},
		{ID: 67, QuestionText: "Как проверить целостность провода?", TestID: test3.ID},
		{ID: 68, QuestionText: "Что такое 'датчик Холла'?", TestID: test3.ID},
		{ID: 69, QuestionText: "Как защитить схему от статического электричества?", TestID: test3.ID},
		{ID: 70, QuestionText: "Какой программой проектируют печатные платы?", TestID: test3.ID},
		{ID: 71, QuestionText: "Для чего нужен ферритовый фильтр на кабеле?", TestID: test3.ID},
		{ID: 72, QuestionText: "Что такое 'LDO-стабилизатор'?", TestID: test3.ID},
	}
	for _, q := range questions3 {
		db.FirstOrCreate(&q, models.Question{ID: q.ID, TestID: q.TestID})
	}

	// Ответы для теста 3
	answers3 := []models.Answer{
		{ID: 241, AnswerText: "Газовая горелка", IsCorrect: false, QuestionID: 65},
		{ID: 242, AnswerText: "Паяльная станция с тонким жалом", IsCorrect: true, QuestionID: 65},
		{ID: 243, AnswerText: "Дрель", IsCorrect: false, QuestionID: 65},
		{ID: 244, AnswerText: "Осциллограф", IsCorrect: false, QuestionID: 65},

		{ID: 245, AnswerText: "Усиливает звук", IsCorrect: false, QuestionID: 66},
		{ID: 246, AnswerText: "Заряжает аккумуляторы", IsCorrect: false, QuestionID: 66},
		{ID: 247, AnswerText: "Переводит сигналы между USB и последовательным портом", IsCorrect: true, QuestionID: 66},
		{ID: 248, AnswerText: "Измеряет температуру", IsCorrect: false, QuestionID: 66},

		{ID: 249, AnswerText: "Осциллографом", IsCorrect: false, QuestionID: 67},
		{ID: 250, AnswerText: "Логическим анализатором", IsCorrect: false, QuestionID: 67},
		{ID: 251, AnswerText: "Мультиметром в режиме прозвонки", IsCorrect: true, QuestionID: 67},
		{ID: 252, AnswerText: "Генератором сигналов", IsCorrect: false, QuestionID: 67},

		{ID: 253, AnswerText: "Датчик температуры", IsCorrect: false, QuestionID: 68},
		{ID: 254, AnswerText: "Датчик влажности", IsCorrect: false, QuestionID: 68},
		{ID: 255, AnswerText: "Датчик звука", IsCorrect: false, QuestionID: 68},
		{ID: 256, AnswerText: "Датчик магнитного поля", IsCorrect: true, QuestionID: 68},

		{ID: 257, AnswerText: "Залить водой", IsCorrect: false, QuestionID: 69},
		{ID: 258, AnswerText: "Нагреть паяльником", IsCorrect: false, QuestionID: 69},
		{ID: 259, AnswerText: "Подключить к сети 220 В", IsCorrect: false, QuestionID: 69},
		{ID: 260, AnswerText: "Использовать антистатический браслет и коврик", IsCorrect: true, QuestionID: 69},

		{ID: 261, AnswerText: "Photoshop", IsCorrect: false, QuestionID: 70},
		{ID: 262, AnswerText: "AutoCAD", IsCorrect: false, QuestionID: 70},
		{ID: 263, AnswerText: "KiCad или Altium Designer", IsCorrect: true, QuestionID: 70},
		{ID: 264, AnswerText: "Microsoft Word", IsCorrect: false, QuestionID: 70},

		{ID: 265, AnswerText: "Усиление сигнала", IsCorrect: false, QuestionID: 71},
		{ID: 266, AnswerText: "Охлаждение проводов", IsCorrect: false, QuestionID: 71},
		{ID: 267, AnswerText: "Измерение напряжения", IsCorrect: false, QuestionID: 71},
		{ID: 268, AnswerText: "Подавление высокочастотных помех", IsCorrect: true, QuestionID: 71},

		{ID: 269, AnswerText: "Генератор шума", IsCorrect: false, QuestionID: 72},
		{ID: 270, AnswerText: "Усилитель мощности", IsCorrect: false, QuestionID: 72},
		{ID: 271, AnswerText: "Стабилизатор напряжения с малым падением", IsCorrect: true, QuestionID: 72},
		{ID: 272, AnswerText: "Датчик освещенности", IsCorrect: false, QuestionID: 72},
	}

	for _, ans := range answers3 {
		db.FirstOrCreate(&ans, models.Answer{ID: ans.ID, QuestionID: ans.QuestionID})
	}
	return nil
}

func SeedTestsAndQuestionsForCourse5(db *gorm.DB) error {
	test1 := models.Test{ID: 10, TestTitle: "Основы микроконтроллеров и среды разработки", CourseID: 5}
	db.FirstOrCreate(&test1)

	// Вопросы для теста 1
	questions1 := []models.Question{
		{ID: 73, QuestionText: "Что такое микроконтроллер?", TestID: test1.ID},
		{ID: 74, QuestionText: "Какая среда разработки чаще всего используется для Arduino?", TestID: test1.ID},
		{ID: 75, QuestionText: "Какой язык программирования используется в Arduino?", TestID: test1.ID},
		{ID: 76, QuestionText: "Что делает функция pinMode() в Arduino?", TestID: test1.ID},
		{ID: 77, QuestionText: "Как загрузить программу в микроконтроллер?", TestID: test1.ID},
		{ID: 78, QuestionText: "Что такое PWM (ШИМ) в микроконтроллерах?", TestID: test1.ID},
		{ID: 79, QuestionText: "Какой сигнал считается аналоговым?", TestID: test1.ID},
		{ID: 80, QuestionText: "Для чего нужна функция delay() в Arduino?", TestID: test1.ID},
	}
	for _, q := range questions1 {
		db.FirstOrCreate(&q, models.Question{ID: q.ID, TestID: q.TestID})
	}

	// Ответы для теста 1
	answers1 := []models.Answer{
		{ID: 273, AnswerText: "Мини-компьютер на одном кристалле с CPU, памятью и периферией", IsCorrect: true, QuestionID: 73},
		{ID: 274, AnswerText: "Аналог транзистора", IsCorrect: false, QuestionID: 73},
		{ID: 275, AnswerText: "Устройство только для обработки аналоговых сигналов", IsCorrect: false, QuestionID: 73},
		{ID: 276, AnswerText: "Датчик температуры", IsCorrect: false, QuestionID: 73},

		{ID: 277, AnswerText: "Visual Studio", IsCorrect: false, QuestionID: 74},
		{ID: 278, AnswerText: "Keil", IsCorrect: false, QuestionID: 74},
		{ID: 279, AnswerText: "Arduino IDE", IsCorrect: true, QuestionID: 74},
		{ID: 280, AnswerText: "MATLAB", IsCorrect: false, QuestionID: 74},

		{ID: 281, AnswerText: "Python", IsCorrect: false, QuestionID: 75},
		{ID: 282, AnswerText: "Java", IsCorrect: false, QuestionID: 75},
		{ID: 283, AnswerText: "C/C++ (упрощенный синтаксис)", IsCorrect: true, QuestionID: 75},
		{ID: 284, AnswerText: "Assembler", IsCorrect: false, QuestionID: 75},

		{ID: 285, AnswerText: "Измеряет напряжение", IsCorrect: false, QuestionID: 76},
		{ID: 286, AnswerText: "Передает данные по UART", IsCorrect: false, QuestionID: 76},
		{ID: 287, AnswerText: "Включает аналоговый сигнал", IsCorrect: false, QuestionID: 76},
		{ID: 288, AnswerText: "Настраивает режим работы пина (вход/выход)", IsCorrect: true, QuestionID: 76},

		{ID: 289, AnswerText: "По Bluetooth", IsCorrect: false, QuestionID: 77},
		{ID: 290, AnswerText: "Через HDMI-порт", IsCorrect: false, QuestionID: 77},
		{ID: 291, AnswerText: "Через программатор или USB-кабель (для Arduino)", IsCorrect: true, QuestionID: 77},
		{ID: 292, AnswerText: "Записать на SD-карту", IsCorrect: false, QuestionID: 77},

		{ID: 293, AnswerText: "Способ измерения сопротивления", IsCorrect: false, QuestionID: 78},
		{ID: 294, AnswerText: "Тип аналогового сигнала", IsCorrect: false, QuestionID: 78},
		{ID: 295, AnswerText: "Протокол связи", IsCorrect: false, QuestionID: 78},
		{ID: 296, AnswerText: "Метод управления мощностью через изменение ширины импульсов", IsCorrect: true, QuestionID: 78},

		{ID: 297, AnswerText: "Только 0 В или 5 В", IsCorrect: false, QuestionID: 79},
		{ID: 298, AnswerText: "Цифровой код", IsCorrect: false, QuestionID: 79},
		{ID: 299, AnswerText: "Сигнал, принимающий любое значение в диапазоне (например, 0–5 В)", IsCorrect: true, QuestionID: 79},
		{ID: 300, AnswerText: "Импульсы фиксированной длины", IsCorrect: false, QuestionID: 79},

		{ID: 301, AnswerText: "Ускорение работы МК", IsCorrect: false, QuestionID: 80},
		{ID: 302, AnswerText: "Чтение аналогового сигнала", IsCorrect: false, QuestionID: 80},
		{ID: 303, AnswerText: "Передача данных", IsCorrect: false, QuestionID: 80},
		{ID: 304, AnswerText: "Приостановка выполнения программы на заданное время (в мс)", IsCorrect: true, QuestionID: 80},
	}

	for _, ans := range answers1 {
		db.FirstOrCreate(&ans, models.Answer{ID: ans.ID, QuestionID: ans.QuestionID})
	}
	test2 := models.Test{ID: 11, TestTitle: "Периферия и протоколы связи", CourseID: 5}
	db.FirstOrCreate(&test2)

	// Вопросы для теста 2
	questions2 := []models.Question{
		{ID: 81, QuestionText: "Какой протокол используется для подключения датчиков температуры (например, DS18B20)?", TestID: test2.ID},
		{ID: 82, QuestionText: "Как подключить дисплей на 1602 к Arduino?", TestID: test2.ID},
		{ID: 83, QuestionText: "Что такое UART?", TestID: test2.ID},
		{ID: 84, QuestionText: "Какой пин в Arduino отвечает за аналоговый ввод?", TestID: test2.ID},
		{ID: 85, QuestionText: "Для чего нужен резистор в 10 кОм при подключении кнопки к Arduino?", TestID: test2.ID},
		{ID: 86, QuestionText: "Какой библиотекой управляют сервоприводом в Arduino?", TestID: test2.ID},
		{ID: 87, QuestionText: "Что делает функция analogRead()?", TestID: test2.ID},
		{ID: 88, QuestionText: "Какой протокол быстрее: I²C или SPI?", TestID: test2.ID},
	}
	for _, q := range questions2 {
		db.FirstOrCreate(&q, models.Question{ID: q.ID, TestID: q.TestID})
	}

	// Ответы для теста 2
	answers2 := []models.Answer{
		{ID: 305, AnswerText: "1-Wire", IsCorrect: true, QuestionID: 81},
		{ID: 306, AnswerText: "I²C", IsCorrect: false, QuestionID: 81},
		{ID: 307, AnswerText: "SPI", IsCorrect: false, QuestionID: 81},
		{ID: 308, AnswerText: "USB", IsCorrect: false, QuestionID: 81},

		{ID: 309, AnswerText: "По Bluetooth", IsCorrect: false, QuestionID: 82},
		{ID: 310, AnswerText: "Через аналоговый вход", IsCorrect: false, QuestionID: 82},
		{ID: 311, AnswerText: "По интерфейсу I²C (с конвертером)", IsCorrect: true, QuestionID: 82},
		{ID: 312, AnswerText: "По SPI", IsCorrect: false, QuestionID: 82},

		{ID: 313, AnswerText: "Последовательный асинхронный протокол связи (RX/TX)", IsCorrect: true, QuestionID: 83},
		{ID: 314, AnswerText: "Параллельный интерфейс", IsCorrect: false, QuestionID: 83},
		{ID: 315, AnswerText: "Беспроводная технология", IsCorrect: false, QuestionID: 83},
		{ID: 316, AnswerText: "Графический интерфейс", IsCorrect: false, QuestionID: 83},

		{ID: 317, AnswerText: "Цифровые пины 0–13", IsCorrect: false, QuestionID: 84},
		{ID: 318, AnswerText: "Только пин 13", IsCorrect: false, QuestionID: 84},
		{ID: 319, AnswerText: "Пины с PWM (~)", IsCorrect: false, QuestionID: 84},
		{ID: 320, AnswerText: "Пин с обозначением \"A0\" (и другие A1-A5)", IsCorrect: true, QuestionID: 84},

		{ID: 321, AnswerText: "Усиление тока", IsCorrect: false, QuestionID: 85},
		{ID: 322, AnswerText: "Защита от перегрева", IsCorrect: false, QuestionID: 85},
		{ID: 323, AnswerText: "Генерация PWM", IsCorrect: false, QuestionID: 85},
		{ID: 324, AnswerText: "Подтяжка к земле (pull-down) для избежания \"плавающего\" сигнала", IsCorrect: true, QuestionID: 85},

		{ID: 325, AnswerText: "Wire.h", IsCorrect: false, QuestionID: 86},
		{ID: 326, AnswerText: "SPI.h", IsCorrect: false, QuestionID: 86},
		{ID: 327, AnswerText: "LiquidCrystal.h", IsCorrect: false, QuestionID: 86},
		{ID: 328, AnswerText: "Servo.h", IsCorrect: true, QuestionID: 86},

		{ID: 329, AnswerText: "Включает цифровой выход", IsCorrect: false, QuestionID: 87},
		{ID: 330, AnswerText: "Передает данные по I²C", IsCorrect: false, QuestionID: 87},
		{ID: 331, AnswerText: "Генерирует ШИМ-сигнал", IsCorrect: false, QuestionID: 87},
		{ID: 332, AnswerText: "Читает аналоговое значение (0–1023) с указанного пина", IsCorrect: true, QuestionID: 87},

		{ID: 333, AnswerText: "I²C", IsCorrect: false, QuestionID: 88},
		{ID: 334, AnswerText: "Они одинаковы", IsCorrect: false, QuestionID: 88},
		{ID: 335, AnswerText: "1-Wire", IsCorrect: false, QuestionID: 88},
		{ID: 336, AnswerText: "SPI (до десятков МГц)", IsCorrect: true, QuestionID: 88},
	}

	for _, ans := range answers2 {
		db.FirstOrCreate(&ans, models.Answer{ID: ans.ID, QuestionID: ans.QuestionID})
	}
	test3 := models.Test{ID: 12, TestTitle: "Продвинутые темы и отладка", CourseID: 5}
	db.FirstOrCreate(&test3)

	// Вопросы для теста 3
	questions3 := []models.Question{
		{ID: 89, QuestionText: "Что такое прерывание (interrupt) в микроконтроллерах?", TestID: test3.ID},
		{ID: 90, QuestionText: "Какой инструмент используют для отладки программ на МК?", TestID: test3.ID},
		{ID: 91, QuestionText: "Что делает директива #define в C/C++?", TestID: test3.ID},
		{ID: 92, QuestionText: "Как уменьшить потребление энергии микроконтроллером?", TestID: test3.ID},
		{ID: 93, QuestionText: "Что такое Watchdog Timer?", TestID: test3.ID},
		{ID: 94, QuestionText: "Как прошить микроконтроллер без программатора (например, через USB)?", TestID: test3.ID},
		{ID: 95, QuestionText: "Какой командой в Arduino проверить наличие данных в Serial-порту?", TestID: test3.ID},
		{ID: 96, QuestionText: "Для чего нужен осциллограф в работе с МК?", TestID: test3.ID},
	}
	for _, q := range questions3 {
		db.FirstOrCreate(&q, models.Question{ID: q.ID, TestID: q.TestID})
	}

	// Ответы для теста 3
	answers3 := []models.Answer{
		{ID: 337, AnswerText: "Событие, которое немедленно прерывает основную программу", IsCorrect: true, QuestionID: 89},
		{ID: 338, AnswerText: "Ошибка в коде", IsCorrect: false, QuestionID: 89},
		{ID: 339, AnswerText: "Увеличение тактовой частоты", IsCorrect: false, QuestionID: 89},
		{ID: 340, AnswerText: "Передача данных по UART", IsCorrect: false, QuestionID: 89},

		{ID: 341, AnswerText: "JTAG/SWD-отладчик (например, ST-Link)", IsCorrect: true, QuestionID: 90},
		{ID: 342, AnswerText: "Мультиметр", IsCorrect: false, QuestionID: 90},
		{ID: 343, AnswerText: "Осциллограф", IsCorrect: false, QuestionID: 90},
		{ID: 344, AnswerText: "Логический анализатор", IsCorrect: false, QuestionID: 90},

		{ID: 345, AnswerText: "Создает макроподстановку (константу)", IsCorrect: true, QuestionID: 91},
		{ID: 346, AnswerText: "Объявляет переменную", IsCorrect: false, QuestionID: 91},
		{ID: 347, AnswerText: "Подключает библиотеку", IsCorrect: false, QuestionID: 91},
		{ID: 348, AnswerText: "Запускает функцию", IsCorrect: false, QuestionID: 91},

		{ID: 349, AnswerText: "Перевести его в режим сна (sleep mode)", IsCorrect: true, QuestionID: 92},
		{ID: 350, AnswerText: "Увеличить тактовую частоту", IsCorrect: false, QuestionID: 92},
		{ID: 351, AnswerText: "Отключить резисторы", IsCorrect: false, QuestionID: 92},
		{ID: 352, AnswerText: "Использовать более толстые провода", IsCorrect: false, QuestionID: 92},

		{ID: 353, AnswerText: "Таймер, который перезагружает МК при зависании", IsCorrect: true, QuestionID: 93},
		{ID: 354, AnswerText: "Датчик температуры", IsCorrect: false, QuestionID: 93},
		{ID: 355, AnswerText: "Генератор случайных чисел", IsCorrect: false, QuestionID: 93},
		{ID: 356, AnswerText: "Инструмент для отладки", IsCorrect: false, QuestionID: 93},

		{ID: 357, AnswerText: "Использовать загрузчик (bootloader)", IsCorrect: true, QuestionID: 94},
		{ID: 358, AnswerText: "Через Bluetooth", IsCorrect: false, QuestionID: 94},
		{ID: 359, AnswerText: "Записать код в EEPROM", IsCorrect: false, QuestionID: 94},
		{ID: 360, AnswerText: "Подключить к HDMI", IsCorrect: false, QuestionID: 94},

		{ID: 361, AnswerText: "Serial.available()", IsCorrect: true, QuestionID: 95},
		{ID: 362, AnswerText: "Serial.read()", IsCorrect: false, QuestionID: 95},
		{ID: 363, AnswerText: "Serial.println()", IsCorrect: false, QuestionID: 95},
		{ID: 364, AnswerText: "Serial.begin()", IsCorrect: false, QuestionID: 95},

		{ID: 365, AnswerText: "Анализ формы сигналов (например, PWM, UART)", IsCorrect: true, QuestionID: 96},
		{ID: 366, AnswerText: "Измерение сопротивления", IsCorrect: false, QuestionID: 96},
		{ID: 367, AnswerText: "Программирование кода", IsCorrect: false, QuestionID: 96},
		{ID: 368, AnswerText: "Зарядка аккумуляторов", IsCorrect: false, QuestionID: 96},
	}

	for _, ans := range answers3 {
		db.FirstOrCreate(&ans, models.Answer{ID: ans.ID, QuestionID: ans.QuestionID})
	}

	return nil

}
