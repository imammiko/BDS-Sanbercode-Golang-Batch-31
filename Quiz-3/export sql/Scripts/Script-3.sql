CREATE TABLE admin_users_groups(
  user_id INT ,
  group_id INT ,
   PRIMARY KEY(user_id,group_id),
   CONSTRAINT fk_user
      FOREIGN KEY(user_id) 
	  REFERENCES admin_users(id)
	  ON DELETE cascade,
	CONSTRAINT fk_group
      FOREIGN KEY(group_id) 
	  REFERENCES admin_groups(id)
	  ON DELETE CASCADE
);