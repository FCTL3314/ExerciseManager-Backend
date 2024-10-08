package domain

type UserUsecase interface {
	GetById(id int64) (*User, error)
	Get(params *FilterParams) (*User, error)
	List(params *Params) (*PaginatedResult[*User], error)
	Create(createUserRequest *CreateUserRequest) (*User, error)
	Login(loginUserRequest *LoginUserRequest) (*TokensResponse, error)
	RefreshTokens(refreshTokenRequest *RefreshTokenRequest) (*TokensResponse, error)
	Update(authUserId, id int64, updateUser *UpdateUserRequest) (*User, error)
	Delete(authUserId, id int64) error
}

type WorkoutUsecase interface {
	GetById(id int64) (*Workout, error)
	Get(params *FilterParams) (*Workout, error)
	List(params *Params) (*PaginatedResult[*Workout], error)
	Create(authUserId int64, createWorkoutRequest *CreateWorkoutRequest) (*Workout, error)
	AddExercise(authUserId, workoutId int64, addExerciseRequest *AddExerciseToWorkoutRequest) (*Workout, error)
	UpdateExercise(authUserId, workoutId, workoutExerciseId int64, updateWorkoutExerciseRequest *UpdateWorkoutExerciseRequest) (*Workout, error)
	RemoveExercise(authUserId, workoutId, workoutExerciseId int64) (*Workout, error)
	Update(authUserId, id int64, updateWorkoutRequest *UpdateWorkoutRequest) (*Workout, error)
	Delete(authUserId, id int64) error
}

type ExerciseUsecase interface {
	GetById(id int64) (*Exercise, error)
	Get(params *FilterParams) (*Exercise, error)
	List(params *Params) (*PaginatedResult[*Exercise], error)
	Create(authUserId int64, createExerciseRequest *CreateExerciseRequest) (*Exercise, error)
	Update(authUserId, id int64, updateExerciseRequest *UpdateExerciseRequest) (*Exercise, error)
	Delete(authUserId, id int64) error
}
