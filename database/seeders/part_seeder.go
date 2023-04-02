package seeders

import (
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/gofiber/fiber/v2"
)

type PartSeeder struct {
	repository        *repositories.PartRepository
	lectureRepository *repositories.LectureRepository
}

// NewPartSeeder instantiates a new PartSeeder
func NewPartSeeder(
	repository *repositories.PartRepository,
	lectureRepository *repositories.LectureRepository,
) *PartSeeder {
	return &PartSeeder{
		repository:        repository,
		lectureRepository: lectureRepository,
	}
}

// Run runs the seeder
func (seeder *PartSeeder) Run(c *fiber.Ctx) (err error) {
	/*
		syncronizer := synch.NewSyncronizer()
		defer syncronizer.Close()
		// Get all lectures
		lectures, err := seeder.lectureRepository.All(c)
		if err != nil {
			return
		}
		// groupedLectures is used to store groups of lectures by course id
		var groupedLectures map[string][]models.Lecture
		for _, lecture := range lectures {
			syncronizer.WG().Add(1)
			go func(lecture models.Lecture) {
				defer syncronizer.WG().Done()
				syncronizer.M().Lock()
				groupedLectures[lecture.CourseId.String()] = append(
					groupedLectures[lecture.CourseId.String()],
					lecture,
				)
				syncronizer.M().Unlock()
			}(lecture)
		}
		syncronizer.WG().Wait()
		// Seed
		var parts []*models.Part
		for _, lectures := range groupedLectures {
			// Sort lectures by lecture.Order
			sort.SliceStable(lectures, func(i, j int) bool {
				return lectures[i].Order < lectures[j].Order
			})
			// chunck the sorted lectures
			chunkLectures := sliceh.Chunk(lectures, (len(lectures)/rand.Intn(5))+1)
			for _, lectures := range chunkLectures {
				for i, lecture := range lectures {
					syncronizer.WG().Add(1)
					go func(lecture models.Lecture) {
						defer syncronizer.WG().Done()
						part := factories.FakePart()
						part.PartableType = reflecth.GetTypeName(lecture)
						part.PartableId = lecture.Id
						part.Name = lecture.Name
						part.Part = (i + 1)
						syncronizer.M().Lock()
						parts = append(parts, &part)
						syncronizer.M().Unlock()
					}(lecture)
				}
			}
		}
		syncronizer.WG().Wait()
		_, err = seeder.repository.Insert(c, parts...)

	*/
	return
}
