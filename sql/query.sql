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
                            JOIN (
                                SELECT
                                    `group_users`.`group_id`
                                FROM
                                    `group_users`
                                    JOIN `users` AS `t0` ON `group_users`.`user_id` = `t0`.`id`
                                WHERE
                                    `t0`.`id` = 10
                            ) AS `t1` ON `groups`.`id` = `t1`.`group_id`
                    ) AS `t1` ON `group_users`.`group_id` = `t1`.`id`
            ) AS `t1` ON `users`.`id` = `t1`.`user_id`
    ) AS `t1` ON `cars`.`owner_id` = `t1`.`id`
WHERE
    NOT (`cars`.`model` = 'Mazda');