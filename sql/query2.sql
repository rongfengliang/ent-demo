SELECT
    DISTINCT `cars`.`id`,
    `cars`.`model`,
    `cars`.`registered_at`
FROM
    `cars`
    JOIN (
        SELECT
            `users`.`id`
        FROM
            `users`
            JOIN (
                SELECT
                    `group_users`.`user_id`
                FROM
                    `group_users`
                    JOIN (
                        SELECT
                            `groups`.`id`
                        FROM
                            `groups`
                        WHERE
                            `groups`.`name` = 'GitHub'
                    ) AS `t1` ON `group_users`.`group_id` = `t1`.`id`
            ) AS `t1` ON `users`.`id` = `t1`.`user_id`
    ) AS `t1` ON `cars`.`owner_id` = `t1`.`id`