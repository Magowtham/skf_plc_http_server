package repository

import (
	"database/sql"
	"fmt"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/entity"
)

type PostgresRepository struct {
	database *sql.DB
}

func NewPostgresRepository(db *sql.DB) PostgresRepository {
	return PostgresRepository{
		database: db,
	}
}

func (repo *PostgresRepository) Init() error {
	query1 := `CREATE TABLE IF NOT EXISTS users (
				user_id VARCHAR(255) PRIMARY KEY,
				email VARCHAR(50) NOT NULL UNIQUE,
				password VARCHAR(255) NOT NULL,
				label VARCHAR(1000) NOT NULL,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);`

	query2 := `CREATE TABLE IF NOT EXISTS admins (
				admin_id VARCHAR(255) PRIMARY KEY,
				email VARCHAR(50) NOT NULL UNIQUE,
				password VARCHAR(255) NOT NULL,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);`

	query3 := `CREATE TABLE IF NOT EXISTS plcs (
				plc_id VARCHAR(255) PRIMARY KEY ,
				user_id VARCHAR(255) NOT NULL,
				label VARCHAR(1000) NOT NULL,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
			);`

	query4 := `CREATE TABLE IF NOT EXISTS driers (
				drier_id VARCHAR(255) PRIMARY KEY,
				plc_id VARCHAR(255) NOT NULL,
				recipe_step_count INTEGER DEFAULT 0,
				label VARCHAR(1000) NOT NULL,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (plc_id) REFERENCES plcs(plc_id) ON DELETE CASCADE
			);`

	query5 := `CREATE TABLE IF NOT EXISTS batches (
				drier_id VARCHAR(255) NOT NULL,
				recipe_step VARCHAR(20) NOT NULL,
				time VARCHAR(20) NOT NULL,
				temp VARCHAR(20) NOT NULL,
				pid VARCHAR(20) NOT NULL,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (drier_id) REFERENCES driers(drier_id) ON DELETE CASCADE
			);`

	query6 := `CREATE TABLE IF NOT EXISTS register_types (
				type VARCHAR(20) PRIMARY KEY ,
				label VARCHAR(1000) NOT NULL,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			);`

	query7 := `CREATE OR REPLACE FUNCTION delete_user(fn_user_id VARCHAR(255))RETURNS VOID AS $$
			    DECLARE
					plc_row RECORD;
				BEGIN
    			FOR plc_row IN
        			SELECT plc_id FROM plcs WHERE user_id = fn_user_id
           		LOOP
             		EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(plc_row.plc_id) || ' CASCADE';
               	END LOOP;


                	DELETE FROM users WHERE user_id = fn_user_id;
                END; $$ LANGUAGE plpgsql;`

	tx, error := repo.database.Begin()

	if error != nil {
		return error
	}

	if _, error := tx.Exec(query1); error != nil {
		tx.Rollback()
		return error
	}

	if _, error := tx.Exec(query2); error != nil {
		tx.Rollback()
		return error
	}

	if _, error := tx.Exec(query3); error != nil {
		tx.Rollback()
		return error
	}

	if _, error := tx.Exec(query4); error != nil {
		tx.Rollback()
		return error
	}

	if _, error := tx.Exec(query5); error != nil {
		tx.Rollback()
		return error
	}

	if _, error := tx.Exec(query6); error != nil {
		tx.Rollback()
		return error
	}

	if _, error := tx.Exec(query7); error != nil {
		tx.Rollback()
		return error
	}

	if error := tx.Commit(); error != nil {
		return error
	}

	return nil
}

func (repo *PostgresRepository) CheckAdminEmailExists(email string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM admins WHERE email = $1)`

	error := repo.database.QueryRow(query, email).Scan(&exists)

	return exists, error
}

func (repo *PostgresRepository) CheckAdminIdExists(adminId string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM admins WHERE admin_id = $1)`

	error := repo.database.QueryRow(query, adminId).Scan(&exists)

	return exists, error
}

func (repo *PostgresRepository) CreateAdmin(admin entity.Admin) error {
	query := `INSERT INTO admins (admin_id,email,password) VALUES ($1,$2,$3)`
	_, error := repo.database.Exec(query, admin.AdminId, admin.Email, admin.Password)
	return error
}

func (repo *PostgresRepository) DeleteAdmin(adminId string) error {
	query := `DELETE FROM admins WHERE admin_id = $1`
	_, error := repo.database.Exec(query, adminId)
	return error
}

func (repo *PostgresRepository) GetAdminByEmail(email string) (entity.Admin, error) {
	var admin entity.Admin

	query := `SELECT admin_id,email,password FROM admins WHERE email = $1;`

	row := repo.database.QueryRow(query, email)

	error := row.Scan(&admin.AdminId, &admin.Email, &admin.Password)

	if error != nil {
		return admin, error
	}

	return admin, nil
}

func (repo *PostgresRepository) CheckUserIdExists(userId string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM users WHERE user_id = $1)`

	error := repo.database.QueryRow(query, userId).Scan(&exists)

	return exists, error
}

func (repo *PostgresRepository) CheckUserEmailExists(email string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`

	error := repo.database.QueryRow(query, email).Scan(&exists)

	return exists, error
}

func (repo *PostgresRepository) CreateUser(user entity.User) error {
	query := `INSERT INTO users (user_id,label,email,password) VALUES ($1,$2,$3,$4)`
	_, error := repo.database.Exec(query, user.UserId, user.Label, user.Email, user.Password)
	return error
}

func (repo *PostgresRepository) DeleteUser(userId string) error {

	query := `SELECT delete_user($1)`

	_, error := repo.database.Exec(query, userId)
	fmt.Println(error)
	return error
}

func (repo *PostgresRepository) GetAllUsers() ([]entity.User, error) {

	var users []entity.User
	var user entity.User

	query := `SELECT user_id,email,label FROM users`

	rows, error := repo.database.Query(query)

	if error != nil {
		return nil, error
	}

	for rows.Next() {
		error := rows.Scan(&user.UserId, &user.Email, &user.Label)

		if error != nil {
			return nil, error
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo *PostgresRepository) GetUserById(userId string) (entity.User, error) {
	var user entity.User

	query := `SELECT email,password FROM users WHERE user_id = $1`

	row := repo.database.QueryRow(query, userId)

	error := row.Scan(&user.Email, &user.Password)

	if error != nil {
		return user, error
	}

	return user, nil
}

func (repo *PostgresRepository) GetUserByEmail(email string) (entity.User, error) {

	var user entity.User

	query := `SELECT user_id,email,label,password FROM users WHERE email=$1`
	row := repo.database.QueryRow(query, email)
	error := row.Scan(&user.UserId, &user.Email, &user.Label, &user.Password)

	if error != nil {
		return user, error
	}

	return user, nil
}

func (repo *PostgresRepository) CheckPlcIdExists(plcId string) (bool, error) {

	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM plcs WHERE plc_id = $1);`

	error := repo.database.QueryRow(query, plcId).Scan(&exists)

	return exists, error
}

func (repo *PostgresRepository) CreatePlc(plc entity.Plc) error {
	query1 := fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %v (
				reg_address VARCHAR(20) PRIMARY KEY,
				reg_type VARCHAR(20) NOT NULL,
				drier_id VARCHAR(255) NOT NULL,
				label VARCHAR(1000) NOT NULL,
				value VARCHAR(20) NOT NULL,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				last_update_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
				FOREIGN KEY (drier_id) REFERENCES driers(drier_id) ON DELETE CASCADE
			)
			`, plc.PlcId)

	query2 := `INSERT INTO plcs (plc_id,user_id,label) VALUES ($1,$2,$3)`

	tx, error := repo.database.Begin()

	if error != nil {
		return error
	}

	_, error = tx.Exec(query1)

	if error != nil {
		tx.Rollback()
		return error
	}

	_, error = tx.Exec(query2, plc.PlcId, plc.UserId, plc.Label)

	if error != nil {
		tx.Rollback()
		return error
	}

	if error := tx.Commit(); error != nil {
		return error
	}

	return nil
}

func (repo *PostgresRepository) DeletePlc(plcId string) error {
	query1 := `DELETE FROM plcs WHERE plc_id = $1;`
	query2 := fmt.Sprintf(`DROP TABLE %s`, plcId)

	tx, error := repo.database.Begin()

	if error != nil {
		return error
	}
	if _, error := tx.Exec(query1, plcId); error != nil {
		tx.Rollback()
		return error
	}

	if _, error := tx.Exec(query2); error != nil {
		tx.Rollback()
		return error
	}

	if error := tx.Commit(); error != nil {
		return error
	}

	return nil
}

func (repo *PostgresRepository) GetPlcsByUserId(userId string) ([]entity.Plc, error) {

	var plcs []entity.Plc
	var plc entity.Plc

	query := `SELECT plc_id,user_id,label FROM plcs WHERE user_id = $1`

	rows, error := repo.database.Query(query, userId)

	if error != nil {
		return nil, error
	}

	for rows.Next() {
		error := rows.Scan(&plc.PlcId, &plc.UserId, &plc.Label)

		if error != nil {
			return nil, error
		}

		plcs = append(plcs, plc)
	}

	return plcs, nil
}

func (repo *PostgresRepository) CreateDrier(drier *entity.Drier) error {
	query := `INSERT INTO driers (drier_id,plc_id,label) VALUES ($1,$2,$3)`
	_, error := repo.database.Exec(query, drier.DrierId, drier.PlcId, drier.Label)
	return error
}

func (repo *PostgresRepository) CheckDrierIdExists(drierId string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM driers WHERE drier_id=$1)`
	error := repo.database.QueryRow(query, drierId).Scan(&exists)
	return exists, error
}

func (repo *PostgresRepository) DeleteDrier(drierId string) error {
	query := `DELETE FROM driers WHERE drier_id = $1`
	_, error := repo.database.Exec(query, drierId)
	return error
}

func (repo *PostgresRepository) GetDriersByUserId(userId string) ([]entity.Drier, error) {

	var driers []entity.Drier
	var direr entity.Drier

	query := `SELECT d.drier_id, d.plc_id, d.recipe_step_count,d.label
				FROM driers d
				JOIN plcs p ON d.plc_id = p.plc_id
				JOIN users u ON p.user_id = u.user_id
				WHERE u.user_id = $1;
			`

	rows, error := repo.database.Query(query, userId)

	if error != nil {
		return nil, error
	}

	for rows.Next() {
		error := rows.Scan(&direr.DrierId, &direr.PlcId, &direr.RecipeStepCount, &direr.Label)
		if error != nil {
			return nil, error
		}

		driers = append(driers, direr)
	}

	return driers, nil
}

func (repo *PostgresRepository) GetDriersByPlcId(plcId string) ([]entity.Drier, error) {

	var driers []entity.Drier
	var direr entity.Drier

	query := `SELECT drier_id,plc_id,recipe_step_count,label FROM driers WHERE plc_id = $1`

	rows, error := repo.database.Query(query, plcId)

	if error != nil {
		return nil, error
	}

	for rows.Next() {
		error := rows.Scan(&direr.DrierId, &direr.PlcId, &direr.RecipeStepCount, &direr.Label)
		if error != nil {
			return nil, error
		}

		driers = append(driers, direr)
	}

	return driers, nil
}

func (repo *PostgresRepository) CheckRegisterAddressExists(plcId string, registerAddress string) (bool, error) {
	var exists bool

	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE reg_address = $1)`, plcId)

	error := repo.database.QueryRow(query, registerAddress).Scan(&exists)

	return exists, error
}

func (repo *PostgresRepository) CheckRegisterAddressAndRegisterTypeExists(plcId string, registerAddress string, registerType string) (bool, error) {
	var exists bool

	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE reg_address = $1 AND reg_type = $2)`, plcId)

	error := repo.database.QueryRow(query, registerAddress, registerType).Scan(&exists)

	return exists, error
}

func (repo *PostgresRepository) CheckRegisterTypeExists(plcId string, drierId string, registerType string) (bool, error) {
	var exists bool

	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE drier_id = $1 AND reg_type = $2)`, plcId)

	error := repo.database.QueryRow(query, drierId, registerType).Scan(&exists)

	return exists, error
}

func (repo *PostgresRepository) UpdateDrierRecipeStepCountAndCreateRegister(plcId string, register *entity.Register) error {
	query1 := `UPDATE driers SET recipe_step_count = recipe_step_count + 1 WHERE drier_id = $1`
	query2 := fmt.Sprintf(`INSERT INTO %v (reg_address,reg_type,drier_id,label,value,last_update_timestamp) VALUES ($1,$2,$3,$4,$5,$6)`, plcId)

	tx, error := repo.database.Begin()

	if error != nil {
		return error
	}

	_, error = tx.Exec(query1, register.DrierId)

	if error != nil {
		tx.Rollback()
		return error
	}

	_, error = tx.Exec(query2, register.RegAddress, register.RegType, register.DrierId, register.Label, register.Value, register.LastUpdateTimestamp)

	if error != nil {
		tx.Rollback()
		return error
	}

	error = tx.Commit()

	return error
}

func (repo *PostgresRepository) CreateRegister(plcId string, register *entity.Register) error {
	query := fmt.Sprintf(`INSERT INTO %v (reg_address,reg_type,drier_id,label,value,last_update_timestamp) VALUES ($1,$2,$3,$4,$5,$6)`, plcId)
	_, error := repo.database.Exec(query, register.RegAddress, register.RegType, register.DrierId, register.Label, register.Value, register.LastUpdateTimestamp)
	return error
}

func (repo *PostgresRepository) UpdateDrierRecipeStepCountAndDeleteRegisterByRegAddress(plcId string, drierId string, registerAddress string) error {
	query1 := `UPDATE driers SET recipe_step_count = recipe_step_count - 1 WHERE drier_id = $1 AND recipe_step_count > 0`
	query2 := fmt.Sprintf(`DELETE FROM %s WHERE reg_address = $1`, plcId)

	tx, error := repo.database.Begin()

	if error != nil {
		return error
	}

	if _, error := tx.Exec(query1, drierId); error != nil {
		tx.Rollback()
		return error
	}

	if _, error := tx.Exec(query2, registerAddress); error != nil {
		tx.Rollback()
		return error
	}

	error = tx.Commit()

	return error
}

func (repo *PostgresRepository) DeleteRegisterByRegAddress(plcId string, registerAddress string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE reg_address = $1`, plcId)

	_, error := repo.database.Exec(query, registerAddress)

	return error
}

func (repo *PostgresRepository) GetRegistersByDrierId(plcId string, drierId string) ([]entity.Register, error) {

	var registers []entity.Register
	var register entity.Register

	query := fmt.Sprintf(`SELECT reg_address,reg_type,drier_id,label,value,last_update_timestamp FROM %s WHERE drier_id = $1`, plcId)

	rows, error := repo.database.Query(query, drierId)

	if error != nil {
		return nil, error
	}

	for rows.Next() {
		error := rows.Scan(&register.RegAddress, &register.RegType, &register.DrierId, &register.Label, &register.Value, &register.LastUpdateTimestamp)

		if error != nil {
			return nil, error
		}

		registers = append(registers, register)
	}

	return registers, nil
}

func (repo *PostgresRepository) CheckRegTypeNameExistsInRegTypes(regTypeName string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS ( SELECT 1 FROM register_types WHERE type = $1 )`

	error := repo.database.QueryRow(query, regTypeName).Scan(&exists)

	return exists, error
}

func (repo *PostgresRepository) CreateRegType(regType *entity.RegisterType) error {
	query := `INSERT INTO register_types (type,label) VALUES ($1,$2)`
	_, error := repo.database.Exec(query, regType.Type, regType.Label)
	return error
}

func (repo *PostgresRepository) DeleteRegType(regTypeName string) error {
	query := `DELETE FROM register_types WHERE type = $1`
	_, error := repo.database.Exec(query, regTypeName)
	return error
}

func (repo *PostgresRepository) GetAllRegisterTypes() ([]entity.RegisterType, error) {
	query := `SELECT type,label FROM register_types`
	var regTypes []entity.RegisterType
	var regType entity.RegisterType

	rows, error := repo.database.Query(query)

	if error != nil {
		return nil, error
	}

	for rows.Next() {
		if error := rows.Scan(&regType.Type, &regType.Label); error != nil {
			return nil, error
		}
		regTypes = append(regTypes, regType)
	}

	return regTypes, error
}

func (repo *PostgresRepository) GetRegisterTypesFromPlcByDrierId(plcId string, drierId string) ([]string, error) {
	query := fmt.Sprintf(`SELECT reg_type FROM %s WHERE drier_id = $1`, plcId)
	var regTypes []string
	var regType string

	rows, error := repo.database.Query(query, drierId)

	if error != nil {
		return nil, error
	}

	for rows.Next() {
		if error := rows.Scan(&regType); error != nil {
			return nil, error
		}

		regTypes = append(regTypes, regType)
	}

	return regTypes, nil
}

func (repo *PostgresRepository) GetRecipeStepCount(drierId string) (int, error) {
	var recipeStepCount int
	query := `SELECT recipe_step_count FROM driers WHERE drier_id = $1`
	error := repo.database.QueryRow(query, drierId).Scan(&recipeStepCount)
	return recipeStepCount, error
}
