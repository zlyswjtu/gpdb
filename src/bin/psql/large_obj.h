/*
 * psql - the PostgreSQL interactive terminal
 *
 * Copyright (c) 2000-2004, PostgreSQL Global Development Group
 *
 * $PostgreSQL: pgsql/src/bin/psql/large_obj.h,v 1.14 2004/08/29 04:13:02 momjian Exp $
 */
#ifndef LARGE_OBJ_H
#define LARGE_OBJ_H

bool		do_lo_export(const char *loid_arg, const char *filename_arg);
bool		do_lo_import(const char *filename_arg, const char *comment_arg);
bool		do_lo_unlink(const char *loid_arg);
bool		do_lo_list(void);

#endif   /* LARGE_OBJ_H */
