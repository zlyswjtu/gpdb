/*-------------------------------------------------------------------------
 *
 * nodeMaterial.h
 *
 *
 *
 * Portions Copyright (c) 1996-2004, PostgreSQL Global Development Group
 * Portions Copyright (c) 1994, Regents of the University of California
 *
 * $PostgreSQL: pgsql/src/include/executor/nodeMaterial.h,v 1.22 2004/08/29 04:13:06 momjian Exp $
 *
 *-------------------------------------------------------------------------
 */
#ifndef NODEMATERIAL_H
#define NODEMATERIAL_H

#include "nodes/execnodes.h"

extern int	ExecCountSlotsMaterial(Material *node);
extern MaterialState *ExecInitMaterial(Material *node, EState *estate);
extern TupleTableSlot *ExecMaterial(MaterialState *node);
extern void ExecEndMaterial(MaterialState *node);
extern void ExecMaterialMarkPos(MaterialState *node);
extern void ExecMaterialRestrPos(MaterialState *node);
extern void ExecMaterialReScan(MaterialState *node, ExprContext *exprCtxt);

#endif   /* NODEMATERIAL_H */
