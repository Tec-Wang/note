package diary

import (
	"database/sql"
	"wangzheng/brain/config"
	"wangzheng/brain/internal/entity"
)

// MysqlDiaryRepository implements Repository using MySQL
type MysqlDiaryRepository struct {
	db *sql.DB
}

func newMySQLDiaryRepository(db *sql.DB) *MysqlDiaryRepository {
	return &MysqlDiaryRepository{db: db}
}

func (r *MysqlDiaryRepository) Save(entry entity.Diary) error {
	stmt, err := r.db.Prepare("INSERT INTO diaries (user_id, content, created_at, updated_at) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(entry.UserID, entry.Content, entry.CreatedAt, entry.UpdatedAt)
	return err
}

func (r *MysqlDiaryRepository) GetAll(userID string) ([]entity.Diary, error) {
	rows, err := r.db.Query("SELECT id, user_id, content, created_at, updated_at FROM diaries WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var entries []entity.Diary
	for rows.Next() {
		var entry entity.Diary
		err = rows.Scan(&entry.ID, &entry.UserID, &entry.Content, &entry.CreatedAt, &entry.UpdatedAt)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func (r *MysqlDiaryRepository) Get(id string) (*entity.Diary, error) {
	stmt, err := r.db.Prepare("SELECT id, user_id, content, created_at, updated_at FROM diaries WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var entry entity.Diary
	err = stmt.QueryRow(id).Scan(&entry.ID, &entry.UserID, &entry.Content, &entry.CreatedAt, &entry.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func (r *MysqlDiaryRepository) Update(entry entity.Diary) error {
	stmt, err := r.db.Prepare("UPDATE diaries SET user_id = ?, content = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(entry.UserID, entry.Content, entry.UpdatedAt, entry.ID)
	return err
}

func (r *MysqlDiaryRepository) Delete(id string) error {
	stmt, err := r.db.Prepare("DELETE FROM diaries WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}

func GetMySQL() *sql.DB {
	sqlConfig := config.MySQLConfig()
	db, err := sql.Open(sqlConfig.Database, sqlConfig.ConnectionString())
	if err != nil {
		panic(err)
	}
	return db
}
