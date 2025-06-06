package seeders

import (
	"Printer3DCourses/internal/models"
	"gorm.io/gorm"
)

func SeedCourses(db *gorm.DB) error {
	authorName := "Дик Роман Владимирович"
	authorTgUsername := "DikRomanV"

	courses := []models.Course{
		{ID: 1, CourseTitle: "Что такое 3D-печать?", ShortDescription: "3D-печать — послойное создание объектов по 3D-модели с минимальными отходами", FullDescription: "3D-печать, или аддитивное производство, представляет собой технологию создания физических объектов путём послойного нанесения материала на основе цифровой 3D-модели. В отличие от традиционных методов, таких как фрезерование или токарная обработка, где материал убирается, здесь он добавляется, что позволяет значительно снизить количество отходов.\n\nЭта технология особенно ценна при производстве изделий со сложной геометрией, которую трудно или невозможно получить другими способами. 3D-печать активно применяется в медицине, авиации, архитектуре и даже в быту, открывая новые возможности в индивидуализированном производстве и прототипировании.", NumberOfParticipants: "150", Duration: 40, Price: 2900, AuthorName: authorName, AuthorTgUsername: authorTgUsername},
		{ID: 2, CourseTitle: "Моделирование в SelfCAD", ShortDescription: "SelfCAD — простой онлайн-сервис для 3D-моделирования и подготовки моделей к печати", FullDescription: "SelfCAD — это удобный и интуитивно понятный онлайн-инструмент для 3D-моделирования, подходящий как для новичков, так и для опытных пользователей. Он позволяет создавать объекты с нуля, используя базовые и сложные формы, а также редактировать их с помощью инструментов деформации, вырезания и добавления текстур.\n\nПосле завершения моделирования пользователи могут экспортировать готовые модели для 3D-печати или других целей. Платформа объединяет в себе возможности моделирования и подготовки к печати, что делает её универсальным решением для творческих и технических проектов.", NumberOfParticipants: "423", Duration: 35, Price: 3200, AuthorName: authorName, AuthorTgUsername: authorTgUsername},
		{ID: 3, CourseTitle: "Настройка 3D-принтера", ShortDescription: "Курс поможет печатать без брака и превратить 3D-принтер в точный инструмент", FullDescription: "Хотите забыть о неудачных распечатках и постоянно ломающихся деталях? Этот курс научит вас печатать идеальные модели с первого раза. Вы узнаете, как правильно настраивать 3D-принтер, подбирать параметры печати и материалы, чтобы результат всегда радовал.\n\nПройдя обучение, вы превратите свой 3D-принтер из капризного устройства в надежный инструмент, готовый воплощать любые ваши идеи. Курс подойдёт как новичкам, так и тем, кто хочет вывести свои навыки 3D-печати на новый уровень.", NumberOfParticipants: "93", Duration: 50, Price: 3500, AuthorName: authorName, AuthorTgUsername: authorTgUsername},
		{ID: 4, CourseTitle: "Радиоэлектроника", ShortDescription: "Курс по радиоэлектронике поможет разобраться в схемах, пайке и создании своих устройств", FullDescription: "Курс по радиоэлектронике — отличный старт для тех, кто хочет понять, как работают электрические схемы и устройства. Вы научитесь читать схемы, подбирать компоненты, паять и собирать собственные проекты — от простых мигалок до полноценной электроники с микроконтроллерами.\n\nМатериал подаётся просто и доступно, даже если у вас нет технического образования. По итогу курса вы сможете не только разбираться в устройствах, но и создавать свои — будь то для хобби, дома или работы.", NumberOfParticipants: "45", Duration: 30, Price: 2700, AuthorName: authorName, AuthorTgUsername: authorTgUsername},
		{ID: 5, CourseTitle: "Программирование микроконтроллеров", ShortDescription: "Курс по микроконтроллерам научит писать код и управлять электроникой на практике", FullDescription: "Курс по программированию микроконтроллеров откроет перед вами мир «умной» электроники. Вы научитесь писать прошивки, подключать датчики, управлять моторами и создавать реальные устройства — от автоматических светильников до мини-роботов.\n\nОбучение проходит на практике: вы пишете код, загружаете его в микроконтроллер и сразу видите результат. Такой подход помогает лучше понять, как «железо» и «софт» работают вместе. Подходит как новичкам, так и тем, кто хочет углубить свои знания в разработке встраиваемых систем.", NumberOfParticipants: "892", Duration: 45, Price: 4100, AuthorName: authorName, AuthorTgUsername: authorTgUsername},
	}

	for _, course := range courses {
		if err := db.FirstOrCreate(&course).Error; err != nil {
			return err
		}
	}

	return nil
}
