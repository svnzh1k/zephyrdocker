package service

func CreateFood(name string, price, categoryID, maxQuantity int) error {
	stmt, err := Database.Prepare("INSERT INTO food (name, price, category_id, max_quantity) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, price, categoryID, maxQuantity)
	return err
}

func UpdateFood(id int, newName string, price, maxQuantity int) error {
	stmt, err := Database.Prepare("UPDATE food SET name = $1, price = $2, max_quantity = $3 WHERE id = $4")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(newName, price, maxQuantity, id)
	return err
}

func RemoveFood(id int) error {
	stmt, err := Database.Prepare("DELETE FROM food WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}

func GetFood() ([]byte, error) {
	// Query for the category tree along with food data in JSON format
	row := Database.QueryRow(`
		WITH RECURSIVE category_tree AS (
			SELECT 
				id, 
				name, 
				parent_id, 
				1 AS level
			FROM categories
			WHERE parent_id IS NULL

			UNION ALL

			SELECT 
				c.id, 
				c.name, 
				c.parent_id, 
				ct.level + 1 AS level
			FROM categories c
			INNER JOIN category_tree ct ON c.parent_id = ct.id
		)

		SELECT json_agg(json_build_object(
			'id', ct.id,
			'name', ct.name,
			'parent_id', ct.parent_id,
			'foods', (
				SELECT json_agg(row_to_json(f))
				FROM food f 
				WHERE f.category_id = ct.id
			),
			'children', (
				SELECT json_agg(json_build_object(
					'id', sub_ct.id,
					'name', sub_ct.name,
					'parent_id', sub_ct.parent_id,
					'level', sub_ct.level,
					'foods', (
						SELECT json_agg(row_to_json(f))
						FROM food f
						WHERE f.category_id = sub_ct.id
					)
				))
				FROM category_tree sub_ct
				WHERE sub_ct.parent_id = ct.id
			)
		)) AS result
		FROM category_tree ct
		WHERE ct.parent_id IS NULL;
	`)

	var jsonResult []byte

	// Scan the JSON result directly from the query result
	err := row.Scan(&jsonResult)
	if err != nil {
		return nil, err
	}

	// Return the JSON result directly
	return jsonResult, nil
}
