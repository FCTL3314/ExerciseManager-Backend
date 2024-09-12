-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS workout_exercises
(
    id          SERIAL PRIMARY KEY,
    workout_id  INT    NOT NULL,
    exercise_id INT    NOT NULL,
    break_time  BIGINT NOT NULL,
    CONSTRAINT fk_workout
        FOREIGN KEY (workout_id)
            REFERENCES workouts (id)
            ON DELETE CASCADE,
    CONSTRAINT fk_exercise
        FOREIGN KEY (exercise_id)
            REFERENCES exercises (id)
            ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS workout_exercises;
-- +goose StatementEnd
